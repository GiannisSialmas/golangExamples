package dummy

import (
	"go-modules/pokemon"
	"log"
)

func Dummy() {

	firstPokemon := pokemon.Get(1)
	log.Println(firstPokemon)
	log.Print("Fetching next pokemon")
	secondPokemon := firstPokemon.Next()
	log.Println(secondPokemon)

}
