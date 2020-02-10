package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("/api/users?legacy=true")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	legacy, legacyOk := q["legacy"]
	if !legacyOk {
		fmt.Println("No legacy parameter found")
	} else {
		fmt.Println("legacy is set to", legacy)
	}

}

# Output 
# > legacy is set to [true]
