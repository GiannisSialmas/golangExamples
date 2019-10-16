package main

import "log"

type Employee struct {
	FirstName string
	LastName  string
}

func (receiver *Employee) UpdateFirstName(newName string) {
	receiver.FirstName = newName
}

func main() {

	giannis := Employee{
		"Giannis",
		"Sialmas",
	}

	log.Printf("First: %s", giannis.FirstName)
	giannis.UpdateFirstName("Basilis")
	log.Printf("First: %s", giannis.FirstName)

}
