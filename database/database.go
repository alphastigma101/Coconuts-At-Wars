/*
This is the database module. It keeps track of the game's current state based on A unique id
It will use types from the game module which the game module includes a module called layout
Which is the whole layout of the game
*/
package database

import (
	//"time"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"

	// My Modules

	"github.com/alphastigma101/Coconuts-At-Wars/game" // Move the interface Properties into layout.go and modify the table struct so it calls them

	// However, The table struct needs to be the parameter of each property
	"github.com/alphastigma101/Coconuts-At-Wars/options"

	// main.go also imports these modules:
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GameTable struct {
	gorm.Model
	Refer uint `gorm:"index:,unique"`
	// Each player can only have one unique setting 1:1
	optionsTable OptionsTable
	// Each player can have many dnd game play-throughs M:M
	dndTable []DndTable `gorm:"foreignKey:Refer;joinForeignKey:DndReferID;References:DndRefer;"`
	// Each player can only have one campaign 1:1
	campaignTable CampaignTable
	// Each player can only be located at one spot a time 1:1
	locationsTable LocationsTable
	// Each player can have multiple weapons
	weaponsTable []WeaponsTable `gorm:"foreignKey:Refer;joinForeignKey:WeaponsReferID;References:WeaponsRefer;"`
	GameData     string         `gorm:"type:TEXT"` // JSON representation of the Game struct
}

type OptionsTable struct {
	gorm.Model
	OptionsID   uint
	OptionsData string `gorm:"type:TEXT"`
}

type LocationsTable struct {
	gorm.Model
	Places         map[string]string
	LocationsRefer uint   `gorm:"index:,unique"`
	LocationsData  string `gorm:"type:TEXT"`
}

type DndTable struct {
	gorm.Model
	DndRefer       uint             `gorm:"index:,unique"`
	locationsTable []LocationsTable `gorm:"many2many:dnd;foreignKey:DndRefer;joinForeignKey:LocationsReferID;References:LocationsRefer"`
	DndData        string           `gorm:"type:TEXT"`
}

type CampaignTable struct {
	gorm.Model
	CampaignRefer  uint             `gorm:"index:,unique"`
	locationsTable []LocationsTable `gorm:"many2many:campaign;foreignKey:CampaignRefer;joinForeignKey:LocationsReferID;References:LocationsRefer"`
	CampaignData   string           `gorm:"type:TEXT"`
}

type WeaponsTable struct {
	gorm.Model
	WeaponsRefer uint `gorm:"index:,unique"`
	weapons      []string
	WeaponsData  string `gorm:"type:TEXT"`
}

func generateKey() uint {
	var b [8]byte // 8 bytes = 64 bits
	_, err := rand.Read(b[:])
	if err != nil {
		panic("Failed to generate random key")
	}
	return uint(binary.BigEndian.Uint32(b[:]))
}

// Creates a file called game.db with GameTable struct
// Each struct is stored via by referencing it inside the Table Struct
func (T *GameTable) Init(table interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	tableData := table.(Layout.Table) // Check to see if it is the Table struct
	gameData := Game.(game.Game)
	tableData.Game = &GameTable{}
	// Check to see if the Game property table has been created
	if !db.Migrator().HasTable(&tableData.Game) {
		tableData.Game = &GameTable{
			Refer: generateKey(),
		}
		// By default, the game will run as 2D
		gameData.Options = &options.Options{
			GameMode: 0,
		}
		// Create the table
		db.AutoMigrate(T)
		gameJSON := Serialize(gameData)
		if gameJSON == "" {
			panic("Failed to serialize game: ")
		}
		GameTable := tableData.Game.(*GameTable)
		GameTable.GameData = gameJSON
		db.Create(GameTable)
	} else {
		GameTable := tableData.Game.(*GameTable) // Check to see if tableData is GameTable type
		// Reference GameTable to modify the variable by retrieving its contents from the database
		result := db.Select([]string{"Refer", "GameData"}).Find(GameTable) // Select TableData, GameData from game
		if result.Error != nil {
			panic("failed to retrieve game tables!")
		}
		Deserialize(GameTable.GameData, &gameData)
	}
	tableData.Options = &OptionsTable{}
	tableData.Campaign = &CampaignTable{}
	tableData.Dnd = &DndTable{}
	tableData.Locations = &LocationsTable{}
	tableData.Weapons = &WeaponsTable{}
	return tableData, gameData
}

// Updates the GameTable struct GameData record
func (T *GameTable) Update(Game interface{}) {
	var existingGame game.Game
	db, err := gorm.Open(sqlite.Open("./database/game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var existingGameTable GameTable
	currentGame := Game.(game.Game)
	Deserialize(existingGameTable.GameData, &existingGame)
	// Need to compare the size of the bytes to see if there was any changes
	// If currentGame bytes are greater than existingGame, then update otherwise do nothing
	db.Model(&existingGameTable).Update("GameData", currentGame)
}

// Inserts values into GameTable struct
func (T *GameTable) Insert(db *gorm.DB) {

}

// Makes queries from the GameTable struct
func (T *GameTable) Query(db *gorm.DB) {

}

// Deletes values from GameTable struct
func (T *GameTable) Delete(db *gorm.DB) {

}

func (T *OptionsTable) Init(optionsTable interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/options.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	tableData := optionsTable.(Layout.Table)
	opts := Game.(options.Options)
	if !db.Migrator().HasTable(&tableData.Options) {
		tableData.Options = &OptionsTable{
			OptionsID: generateKey(),
		}
		optionsTableJSON := Serialize(opts)
		if optionsTableJSON == "" {
			panic("Failed to serialize!")
		}
		db.AutoMigrate(T)
		OptionTable := tableData.Options.(*OptionsTable)
		OptionTable.OptionsData = optionsTableJSON
		db.Create(OptionTable)
	} else {
		OptionTable := tableData.Options.(*OptionsTable) // Check to see if tableData is GameTable type
		// Reference GameTable to modify the variable by retrieving its contents from the database
		result := db.Select([]string{"OptionsID", "OptionsData"}).Find(OptionTable) // Select OptionsID, OptionsData from options
		if result.Error != nil {
			panic("failed to retrieve game tables!")
		}
		Deserialize(OptionTable.OptionsData, &opts)
	}
	return optionsTable, Game
}

// Pass in the options Struct into Options
func (T *OptionsTable) Update(Options interface{}) {
	existingOptions := Options.(options.Options)
	db, err := gorm.Open(sqlite.Open("./database/options.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Deserialize(T.OptionsData, existingOptions)
	db.Model(T).Update("OptionsData", existingOptions)
}

func (T *OptionsTable) Insert(db *gorm.DB) {

}

func (T *OptionsTable) Query(db *gorm.DB) {

}

func (T *OptionsTable) Delete(db *gorm.DB) {

}

func (T *DndTable) Init(currTable interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/dnd.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// tableData := currTable.(Table)
	// tableData.Dnd := &DndTable{
	//DndRefer: generateKey(),
	//}
	db.AutoMigrate(T)
	return currTable, Game
}

func (T *DndTable) Update(Dnd interface{}) {
	existingDnd := Dnd.(dnd.Dnd)
	db, err := gorm.Open(sqlite.Open("./database/dnd.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Deserialize(T.DndData, existingDnd)
	db.Model(T).Update("DndData", existingDnd)
}

func (T *DndTable) Insert(db *gorm.DB) {

}

func (T *DndTable) Query(db *gorm.DB) {

}

func (T *DndTable) Delete(db *gorm.DB) {

}

func (T *CampaignTable) Init(i interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/campaign.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(T)
	return i, Game
}

func (T *CampaignTable) Update(Campaign interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/campaign.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	campaign := Campaign.(game.Game)
	Deserialize(T.CampaignData, campaign.GameState)
	db.Model(T).Update("CampaignData", campaign)
}

func (T *CampaignTable) Insert(db *gorm.DB) {

}

func (T *CampaignTable) Query(db *gorm.DB) {

}

func (T *CampaignTable) Delete(db *gorm.DB) {

}

func (T *WeaponsTable) Init(i interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/weapons.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(T)
	return i, Game
}

// Need to keep track of weapons of each player has picked up
// Player struct inside of dnd.go has a field called id which will be used
// Search through Player.Bag.Weapons
func (T *WeaponsTable) Update(Weapons interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/weapons.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	weapons := Weapons.(game.Game)
	//Deserialize(existingWeaponsTable.WeaponsData, existingWeapons.)
	// Need to compare the size of the bytes to see if there was any changes
	// If currentGame bytes are greater than existingGame, then update otherwise do nothing
	// if campaign.Campaign > existingCampaign.Campaign
	db.Model(T).Update("WeaponsData", weapons)
}

func (T *WeaponsTable) Insert(db *gorm.DB) {

}

func (T *WeaponsTable) Query(db *gorm.DB) {

}

func (T *WeaponsTable) Delete(db *gorm.DB) {

}

func (T *LocationsTable) Init(i interface{}, Game interface{}) (interface{}, interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/locations.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(T)
	return i, Game
}

// Search through the Game.Player.coords and compare the coords of each location avialable
func (T *LocationsTable) Update(Locations interface{}) {
	db, err := gorm.Open(sqlite.Open("./database/locations.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var existingGameTable LocationsTable
	locations := Locations.(game.Game)
	if locations.Options.DndMode == 1 {
		// Pull all the images and render them and update the table
	}
	db.Model(&existingGameTable).Update("LocationData", locations.Players.Bag.Location)
}

func (T *LocationsTable) Insert(db *gorm.DB) {

}

func (T *LocationsTable) Query(db *gorm.DB) {

}

func (T *LocationsTable) Delete(db *gorm.DB) {

}

// SerializeGame converts the Game struct into JSON for storage
func Serialize(g interface{}) string {
	gameJSON, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(gameJSON)
}

// DeserializeGame converts JSON back into a Game struct
// When deserializing, the parameter out must either be deferenced or referenced with a pointer to match
// the serialization
func Deserialize(data string, out interface{}) error {
	return json.Unmarshal([]byte(data), out)
}
