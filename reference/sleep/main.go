package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

    time.Sleep(2 * time.Second)

    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
}