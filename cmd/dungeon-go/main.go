package main

import "fmt"

//Стуктура героя
type Hero struct {
	Name       string
	Health     int
	Mana       int
	Level      int
	Experience int
}

//Метод атаки
func (h *Hero) Attack() int {
	damage := h.Level * 10
	h.Experience += 5
	return damage
}

func (h *Hero) DisplayInfo() {

}

func (h *Hero) Heal() {

}

//Создание нового героя(конструктор)
func NewHero(name string) Hero {
	return Hero{
		Name:       name,
		Health:     100,
		Mana:       100,
		Level:      1,
		Experience: 0,
	}
}

func main() {
	//создание героя
	juggernaut := NewHero("Juggernaut")
	fmt.Println(juggernaut)
	damage := juggernaut.Attack()
	fmt.Printf("%s наносит урон: %d", juggernaut.Name, damage)
	fmt.Println(juggernaut)
}
