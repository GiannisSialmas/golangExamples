package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var listenAdress string

func main() {
	listenAdress = getListenAdress()

	//channels to handle gracefull shutdown
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	server := newWebserver(logger)

	// Start graceful shutdown routine
	go gracefullShutdown(server, logger, quit, done)

	logger.Println("Listening to", listenAdress)
	log.Fatal(http.ListenAndServe(listenAdress, nil))

	// Deadlock until server is stopped
	<-done
	logger.Println("Server stopped")

}

// Calculates the listen adress or the HTTP server from environmental variables
func getListenAdress() string {
	// Get port from env variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		fmt.Printf("No port given. Defaulting to port %s\n", port)
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("Invalid port", port)
	}
	//Get interface to bind to from env variable
	bindInterface := os.Getenv("BIND_INTERFACE")
	var bindAdress string
	if bindInterface == "" {
		bindAdress = "0.0.0.0"
		fmt.Printf("Defaulting to bindInterface %s\n", bindInterface)
	} else {

		ief, err := net.InterfaceByName(bindInterface)

		if err != nil {
			log.Fatal("Error binding interface:", err)
		}
		addrs, err := ief.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		bindAdress = addrs[0].(*net.IPNet).IP.String()
	}
	return fmt.Sprintf("%s:%s", bindAdress, port)
}

// Returns a new webserver instance with k8s probes ready
func newWebserver(logger *log.Logger) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "This is the / path\n")
	})
	// router.HandleFunc("/health/liveness", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	io.WriteString(w, "Healthy\n")
	// })
	// router.HandleFunc("/health/readiness", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	io.WriteString(w, "Ready\n")
	// })
	return &http.Server{
		Addr:     listenAdress,
		Handler:  router,
		ErrorLog: logger,
	}
}

// Try to gracefully shut down the server
func gracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	// quit receives its value from os.Signal, done is a channel we add values to
	// Cause a deadlock until we receive a value from is.Signal
	<-quit
	logger.Println("SIGINT received. Server will try to shut down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prevent new keep alive connections
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
