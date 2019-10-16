package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func (pokemon Pokemon) Next() Pokemon {

	currentPokemonID := pokemon.ID
	log.Printf("Next pokemon id is %d", currentPokemonID+1)
	return Get(currentPokemonID + 1)

}

func Get(id int) Pokemon {

	url := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d/", id)
	fmt.Println("Url is %s", url)
	response, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	var data Pokemon
	if err := json.Unmarshal(responseData, &data); err != nil {
		log.Print(err)
	}
	return data
}

func main() {

	firstPokemon := Get(1)
	log.Println(firstPokemon)
	log.Print("Fetching next pokemon")
	secondPokemon := firstPokemon.Next()
	log.Println(secondPokemon)
}
