package game

import (
	// My modules
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	coop "github.com/alphastigma101/Coconuts-At-Wars/cooperative"
	"github.com/alphastigma101/Coconuts-At-Wars/main_game"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

// Layout of the game. It will copied to the database assigned with a unique player id
// It also will be the options the user will see once they get pass the titlescreen
type Game struct {
	Options     *options.Options // Options that the user can configure
	Campaign    *main_game.MainGame
	Dnd         *dnd.Dnd
	Cooperative *coop.Cooperative
	Players     *dnd.Player
}
