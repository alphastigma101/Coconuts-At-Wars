package main

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
	// My own modules
	database "github.com/alphastigma101/Coconuts-At-Wars/database"
	game "github.com/alphastigma101/Coconuts-At-Wars/game"
	input "github.com/alphastigma101/Coconuts-At-Wars/input_handler" // Might have to move input_handler inside of game.go and make a type alias to it
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

func main() {
	app := game.Game{}
	table := Layout.Table{}
	table.Game = &database.GameTable{}
	tempTable, tempGameTable := table.Game.Init(table, app)
	table = tempTable.(Layout.Table)
	app = tempGameTable.(game.Game)
	// Initialize window first
	r1.InitWindow(800, 450, "Coconuts At Wars")
	r1.SetTargetFPS(60)
	if app.Options.GameMode == 1 {
		updatedOpts, updatedTable := options.UpdateOptions(app.Options, &table)
		app.Options = updatedOpts
		table = updatedTable.(Layout.Table)
		app.Options.Game3D.InitializeTitleScreen() // Initalize the game startup
		app.Options.Game3D.InitializeMainMenuScreen(&app, &table)
	} else {
		updatedOpts, updatedTable := options.UpdateOptions(app.Options, &table)
		app.Options = updatedOpts
		table = updatedTable.(Layout.Table)
		app.Options.Game2D.InitializeTitleScreen() // Initalize the game startup
		app.Options.Game2D.InitializeMainMenuScreen(&app, &table)
	}
	if app.GameActor == nil {
		app.GameActor = &game.Actor{
			Health: 100,
		}
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
	// The method below will update the whole game causing it to sync properly
	//table.Game.Update(&app)                    // Updates when it needs too
}
