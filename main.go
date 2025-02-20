package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// My own modules
	database "github.com/alphastigma101/Coconuts-At-Wars/database"
	game "github.com/alphastigma101/Coconuts-At-Wars/game"
	input "github.com/alphastigma101/Coconuts-At-Wars/input_handler"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

func generateKey() uint64 {
	var b [8]byte // 8 bytes = 64 bits
	_, err := rand.Read(b[:])
	if err != nil {
		panic("Failed to generate random key")
	}
	return binary.BigEndian.Uint64(b[:]) // Convert bytes to uint64
}

func DB(db *gorm.DB, tableData *database.Table, gameData *game.Game) (*database.Table, *game.Game) {
	tableJSON, err := database.Serialize(&tableData)
	if err != nil {
		tableData = &database.Table{
			Game:      &database.GameTable{},
			Options:   &database.DndTable{},
			Campaign:  &database.CampaignTable{},
			Weapons:   &database.WeaponsTable{},
			Locations: &database.LocationsTable{},
			Load:      &database.LoadTable{},
			Save:      &database.SaveTable{},
		}
	} else {
		var existingGameTable database.GameTable
		err := database.Deserialize(existingGameTable.TableData, &tableData)
		if err != nil {
			panic("Failed to deserialize game: " + err.Error())
		}
	}
	if !db.Migrator().HasTable(&tableData.Game) {
		gameData.Options.GameMode = 0
		fmt.Println("GameTable does not exist. Creating new table and game entry.")
		// Create the table
		tableData.Game.Init(db)
		// Create a new Game entry
		gameJSON, err := database.Serialize(&gameData)
		if err != nil {
			panic("Failed to serialize game: " + err.Error())
		}
		newGame := database.GameTable{
			Refer:     generateKey(),
			GameData:  gameJSON,
			TableData: tableJSON,
		}
		db.Create(&newGame)
	} else {
		//fmt.Println("GameTable exists. Initializing existing game entry.")
		var existingGameTable database.GameTable
		err := database.Deserialize(existingGameTable.GameData, &gameData)
		if err != nil {
			panic("Failed to deserialize game: " + err.Error())
		}
	}
	return tableData, gameData
}

func main() {
	app := game.Game{}
	table := database.Table{}
	db, err := gorm.Open(sqlite.Open("./database/game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB(db, &table, &app)
	if app.Options.GameMode == 1 {
		fmt.Println("3D has been enabled by user!")
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game3D.InitializeTitleScreen() // Initalize the game startup
		table.Game.Update(db, &app)
	} else {
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game2D.InitializeTitleScreen() // Initalize the game startup
		table.Game.Update(db, &app)
	}
	// Need a handful of if statements that pull in the start menu of the game
	// It will display the game with three coconuts huddle around
	// Two of them will have army helments while the other coconut on the far right will have
	// A mullet. Everytime the user starts up the game the coconut with the mullet will have different cosmetics
	// Such as googles, a cone on its head, etc
	// They will be huddled around a gray-ish steel circlular object which the landscape will be mud while airplanes fly around
	// Similar to inheritance because an interface needs some object that implements
	// it's properties
	actor := input.GameActor{
		Health: 100,
	}
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
