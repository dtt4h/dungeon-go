package models

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
