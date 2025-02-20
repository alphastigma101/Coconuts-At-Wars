package game

import (
	// My modules
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	"github.com/alphastigma101/Coconuts-At-Wars/main_game"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

// Layout of the game. It will copied to the database assigned with a unique player id
type Game struct {
	Options  *options.Options // Options that the user can configure
	Campaign *main_game.MainGame
	Dnd      *dnd.Dnd
}
