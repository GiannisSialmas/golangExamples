package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 50)
	go count("sheep", c)

	// Way 1: get value from channel and check if open
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// Way 2: iterate the values in the channel
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	// The channel should be closed always by the sender. If the sender sends a message
	// to a closed channel, the program will panic and error out
	close(c)
}
