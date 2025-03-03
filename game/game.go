package game

// This Module holds the Game struct which consists of nested structs that keeps track of the user's data
import (
	// My modules
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	coop "github.com/alphastigma101/Coconuts-At-Wars/cooperative"
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
	"github.com/alphastigma101/Coconuts-At-Wars/main_game"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

type DataBaseProperties Layout.Properties

// During the campaign, if the user clicks a certain button, it will pop up a menu
// They can either save, exit, load, and connect
type GameOptions interface {
	Options() interface{}
}

// Struct that will keep track of the damage and the position of the player
type gameActor struct {
	Health   int
	Position Layout.Position
}

type Actor gameActor

// Layout of the game. It will copied to the database assigned with a unique player id
// It also will be the options the user will see once they get pass the titlescreen
type Game struct {
	Options     *options.Options // Options that the user can configure
	Campaign    *main_game.MainGame
	Dnd         *dnd.Dnd
	Cooperative *coop.Cooperative
	Players     *dnd.Player
	GameActor   *Actor
}

func (g *Game) InitializeCampaignScreen() {
	// Reuse the main menu box to draw out the campaign screen:
	//g.Options.Game2D.(Layout.Game2D).GameLayout.MainMenu.Rectangle
	// Should have Continue, StartOver, Exit
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		// Everything needs to be unloaded such as images, textures, etc
		g.Options.Game2D.InitializeMainMenuScreen(g) // If user clicks exit, go back to main menu screen
	}
	panic("Not Implemented Yet!")
}

func (g *Game) InitializeOptionsScreen() interface{} {
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		g.Options.Game2D.InitializeMainMenuScreen(g) // If user clicks exit, go back to main menu screen
	}
	panic("Not Implemented Yet!")
}
