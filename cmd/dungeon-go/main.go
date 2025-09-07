package main

import (
	"dungeon-go/internal/models" //importing package models
	"fmt"
)

func main() {
	// Creating a hero using the constructor
	juggernaut := models.NewHero("Juggernaut")

	fmt.Println(juggernaut)
}
