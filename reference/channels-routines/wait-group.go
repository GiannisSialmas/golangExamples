package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// You can pass the wait group in the count function. As it's not semantically correct(it means it's not a part of count)
	// we make an unamed function and handle here as we have the context of the wait group
	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait() // This blocks until wait group is 0???
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
