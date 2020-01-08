package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getListenAdress() string {
	if port, ok := os.LookupEnv("PORT"); !ok {
		 log.Fatalf("No port given")
	}
	
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port", port)
	}
	return port
}

func main() {
	port := getListenAdress()
	fmt.Println("Valid Port is:", port)

}
