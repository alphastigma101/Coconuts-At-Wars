package main

import (

	// My own modules
	"fmt"

	"github.com/alphastigma101/Coconuts-At-Wars/input_handler"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

func main() {
	// Game will always run 2D graphics by default
	game := options.Game{
		Mode: 0,
	}
	actor := input_handler.GameActor{
		Health: 100,
	}
	if game.Mode == 1 {
		fmt.Println("3D has been enabled by user!")
		game = *(options.CreateGame(&game))
	}
	// Need a handful of if statements that pull in the start menu of the game
	// It will display the game with three coconuts huddle around
	// Two of them will have army helments while the other coconut on the far right will have
	// A mullet. Everytime the user starts up the game the coconut with the mullet will have different cosmetics
	// Such as googles, a cone on its head, etc
	// They will be huddled around a gray-ish steel circlular object which the landscape will be mud while airplanes fly around
	// By default it will be rendered as 2D graphics
	// Similar to inheritance because an interface needs some object that implements
	// it's properties
	handler := &input_handler.InputHandler{
		ButtonX: &input_handler.JumpCommand{},
		ButtonY: &input_handler.FireCommand{},
		ButtonA: &input_handler.DuckCommand{},
		ButtonB: &input_handler.ReloadCommand{},
	}

	command := input_handler.HandleInput(handler)
	if command != nil {
		command.Execute(&actor)
	}
}
