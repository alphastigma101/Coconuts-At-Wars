package options

import (
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
)

type Game2D Layout.Game2D

// GameMode represents the dimension type (2D or 3D)
type gameMode int
type dndMode int

type Options struct {
	GameMode gameMode
	DndMode  dndMode
	Game2D   Layout.Render
	Game3D   Layout.Render
}

// CreateRenderer factory function to create appropriate renderer
func UpdateOptions(game *Options) *Options {
	if game.GameMode == 1 {
		// Need to update the Options Table
		//game.Game3D = Layout.GetGame3D()
		return game
	} else {
		game.Game2D = Layout.GetGame2D()
		return game
	}
}

func (g Game2D) InitializeOptionsScreen() interface{} {
	// Use GetGame2D to check and see if g is the instance of game2d
	return nil
}
