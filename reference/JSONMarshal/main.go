package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	basicJSONExample()
}

func basicJSONExample() {

	// The properties must be uppercase so that they can be exported and used by the Marshal command
	response := struct {
		Status  bool
		Message string
	}{
		Status:  true,
		Message: "You are correct",
	}
	responseJSON, error := json.Marshal(response)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("%T\n", string(responseJSON))
	fmt.Printf("%+v", string(responseJSON))
}
