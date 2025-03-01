package main_game

import (
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
)

type MainGame struct {
	Players dnd.Player
	Enemies map[string]string // Stores the images that will be either rendered in 2D or 3D
}

type SaveData map[uint]*MainGame // Represents the progress through the game

func (d *MainGame) Campaign() {
	panic("Function has not been implemented yet!")
}

func (d *MainGame) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *MainGame) Locations() {
	panic("Function has not been implemented yet!")
}
