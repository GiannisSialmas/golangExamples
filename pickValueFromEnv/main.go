package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("API KEY:", os.Getenv("API_KEY"))
}
