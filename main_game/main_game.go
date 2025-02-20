package main_game

import (
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
)

type MainGame struct {
	players dnd.Players
	Enemies map[string]string // Stores the images that will be either rendered in 2D or 3D
}

func (d *MainGame) Campaign() {
	panic("Function has not been implemented yet!")
}

func (d *MainGame) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *MainGame) Locations() {
	panic("Function has not been implemented yet!")
}
