package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	//"input_handler/render2d"

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
	tableData.Game = &database.GameTable{}
	if !db.Migrator().HasTable(&tableData.Game) {
		tableData = &database.Table{
			Game: &database.GameTable{
				Refer: generateKey(),
			},
			Options: &database.OptionsTable{
				OptionsID: generateKey(),
			},
			Dnd: &database.DndTable{
				DndRefer: generateKey(),
			},
			Campaign: &database.CampaignTable{
				CampaignRefer: generateKey(),
			},
			Weapons: &database.WeaponsTable{
				WeaponsRefer: generateKey(),
			},
			Locations: &database.LocationsTable{
				LocationsRefer: generateKey(),
			},
			Load: &database.LoadTable{
				LoadID: generateKey(),
			},
			Save: &database.SaveTable{
				SaveRefer: generateKey(),
			},
		}
		tableJSON := database.Serialize(tableData)
		if tableJSON == "" {
			panic("Failed to serialize!")
		}
		gameData.Options = &options.Options{
			GameMode: 0,
		}
		fmt.Println("GameTable does not exist. Creating new table and game entry.")
		// Create the table
		tableData.Game.Init(db)
		// Create a new Game entry
		gameJSON := database.Serialize(gameData)
		if gameJSON == "" {
			panic("Failed to serialize game: ")
		}
		GameTable := tableData.Game.(*database.GameTable)
		GameTable.GameData = gameJSON
		GameTable.TableData = tableJSON
		db.Create(&GameTable)
	} else {
		GameTable := tableData.Game.(*database.GameTable)
		result := db.Find(&GameTable)                   // select * from GameTable;
		currentTableData, ok := result.Get("TableData") // Search for the specific key
		if !ok {
			panic("TableData was not successfully stored!")
		}
		tableData = currentTableData.(*database.Table)
		currentGameData, ok := result.Get("GameData") // Search for the specific key
		if !ok {
			panic("GameData was not successfully stored!")
		}
		gameData = currentGameData.(*game.Game)

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
		table.Game.Update(db, &app)                // Updates when it needs too
	} else {
		opts := *(options.UpdateOptions(app.Options))
		app.Options = &opts
		app.Options.Game2D.InitializeTitleScreen() // Initalize the game startup
		table.Game.Update(db, &app)                // Updates when it needs too
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
