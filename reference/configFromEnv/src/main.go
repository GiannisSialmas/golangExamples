package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getListenAdress() string {
	port := os.Getenv("PORT")
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
