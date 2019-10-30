package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

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

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	listenAdress := getListenAdress()
	fmt.Println("Listening to", listenAdress)
	log.Fatal(http.ListenAndServe(listenAdress, nil))
}
