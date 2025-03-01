package main

import (
	"fmt"
	// My own modules
	database "github.com/alphastigma101/Coconuts-At-Wars/database"
	game "github.com/alphastigma101/Coconuts-At-Wars/game"
	input "github.com/alphastigma101/Coconuts-At-Wars/input_handler"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

func main() {
	app := game.Game{}
	table := database.Table{}
	table.Game = &database.GameTable{}
	tempTable, tempGameTable := table.Game.Init(table, app)
	table = tempTable.(database.Table)
	app = tempGameTable.(game.Game)
	if app.Options.GameMode == 1 {
		table.Options = &database.OptionsTable{}
		//tempOptionTable, tempGameTable := table.Options.Init(db, table, app)
		fmt.Println("3D has been enabled by user!")
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game3D.InitializeTitleScreen() // Initalize the game startup
		table.Game.Update(&app)                    // Updates when it needs too
	} else {
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game2D.InitializeTitleScreen() // Initalize the game startup
		table.Game.Update(&app)                    // Updates when it needs too
	}
	actor := input.GameActor{
		Health: 100,
	}
	// Game Loop
	gameStart := true
	for gameStart != false {
		handler := &input.InputHandler{
			ButtonX: &input.JumpCommand{},
			ButtonY: &input.FireCommand{},
			ButtonA: &input.DuckCommand{},
			ButtonB: &input.ReloadCommand{},
		}

		command := input.HandleInput(handler)
		if command != nil {
			command.Execute(&actor)
		}
	}
}
