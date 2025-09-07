package models

import "fmt"

type Hero struct {
	Name       string
	Health     int
	Mana       int
	Level      int
	Experience int
}

func NewHero(name string) Hero {
	return Hero{
		Name:       name,
		Health:     100,
		Mana:       100,
		Level:      1,
		Experience: 0,
	}
}

// Attack method
func (h *Hero) Attack() int {
	damage := h.Level * 10
	h.Experience += 5
	return damage
}

// Display method
func (h *Hero) DisplayInfo() {
	fmt.Printf(
		"\tname: %s\n\thealth: %d\n\tmana: %d\n\tlevel: %d\n\texperience: %d\n",
		h.Name, h.Health, h.Mana, h.Level, h.Experience,
	)
}

// Heal method
func (h *Hero) Heal() {

}
