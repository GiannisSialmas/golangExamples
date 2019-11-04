//test things here
package main

import (
	"fmt"
	"time"
)

var blockForever chan bool

func main() {
	go func() {
		for now := range time.Tick(time.Second) {
			fmt.Println(now)
		}
	}()

	fmt.Println("Main started")
	<-blockForever

}
