package main

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
	// My own modules
	database "github.com/alphastigma101/Coconuts-At-Wars/database"
	game "github.com/alphastigma101/Coconuts-At-Wars/game"
	input "github.com/alphastigma101/Coconuts-At-Wars/input_handler" // Might have to move input_handler inside of game.go and make a type alias to it
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
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game3D.InitializeTitleScreen() // Initalize the game startup
		//table.Game.Update(&app)                    // Updates when it needs too
		app.Options.Game3D.InitializeMainMenuScreen(&app)
	} else {
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game2D.InitializeTitleScreen() // Initalize the game startup
		//table.Game.Update(&app)                    // Updates when it needs too
		app.Options.Game2D.InitializeMainMenuScreen(&app)
	}
	app.GameActor = &game.Actor{
		Health: 100,
	}
	// Game Loop
	for !r1.WindowShouldClose() {
		handler := &input.InputHandler{
			ButtonX: &input.JumpCommand{},
			ButtonY: &input.FireCommand{},
			ButtonA: &input.DuckCommand{},
			ButtonB: &input.ReloadCommand{},
		}

		command := input.HandleInput(handler)
		if command != nil {
			command.Execute(app.GameActor)
		}
	}
	// TODO: Call in all the database properties and sync the data with it
	// You can't update anything other than here due to circular dependcy
}
