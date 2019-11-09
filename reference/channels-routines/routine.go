package main

import (
	"fmt"
	"time"
)

// Both sheep and fish will be printed in the console as one runs async and the other sync
func main() {
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
