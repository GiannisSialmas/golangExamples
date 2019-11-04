package main

import (

	"log"
	"net/http"
	"fmt"
	"redis-example/rediscache"
)

var blockForever chan bool

// var listenAdress string
func main() {
	fmt.Println("Beginning main. This will never exit")
	fmt.Println("Print:", rediscache.CacheActive)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// // Deadlock to never exit
	// <- blockForever
}


