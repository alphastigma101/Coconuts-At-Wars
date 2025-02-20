package database

import (
	//"time"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	"encoding/json"
)

type GameTable struct {
	gorm.Model
	Refer uint64 `gorm:"index:,unique"`
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
	// Each player can have multiple saved files
	saveTable []SaveTable `gorm:"foreignKey:Refer;joinForeignKey:SaveReferID;References:SaveRefer;"`
	// Each player can load one save file
	loadTable LoadTable
	GameData  string `gorm:"type:TEXT"` // JSON representation of the Game struct
	TableData string `gorm:"type:TEXT"` // JSON representation of the Table struct that holds everything
}

type OptionsTable struct {
	gorm.Model
	OptionsID   uint64
	OptionsData string `gorm:"type:TEXT"`
}

type LocationsTable struct {
	gorm.Model
	Places         map[string]string
	LocationsRefer uint64 `gorm:"index:,unique"`
	LocationsData  string `gorm:"type:TEXT"`
}

type DndTable struct {
	gorm.Model
	DndRefer       uint64           `gorm:"index:,unique"`
	locationsTable []LocationsTable `gorm:"many2many:dnd;foreignKey:DndRefer;joinForeignKey:LocationsReferID;References:LocationsRefer"`
	DndData        string           `gorm:"type:TEXT"`
}

type CampaignTable struct {
	gorm.Model
	CampaignRefer  uint64           `gorm:"index:,unique"`
	locationsTable []LocationsTable `gorm:"many2many:campaign;foreignKey:CampaignRefer;joinForeignKey:LocationsReferID;References:LocationsRefer"`
	CampaignData   string           `gorm:"type:TEXT"`
}

type WeaponsTable struct {
	gorm.Model
	WeaponsRefer uint64 `gorm:"index:,unique"`
	weapons      []string
	WeaponsData  string `gorm:"type:TEXT"`
}

type LoadTable struct {
	gorm.Model
	Load     map[string]string // key will be the time and date the file was loaded by the user and value will be the saved file to load
	LoadID   uint64
	LoadData string `gorm:"type:TEXT"`
}

type SaveTable struct {
	gorm.Model
	Save      map[string]string // key will be the time and dat the user saved their data and the value will be the saved file to save
	SaveRefer uint64            `gorm:"index:,unique"`
	SaveData  string            `gorm:"type:TEXT"`
}

type Properties interface {
	Init(db *gorm.DB)
	Update(db *gorm.DB, game interface{})
	Insert(db *gorm.DB)
	Query(db *gorm.DB)
	Delete(db *gorm.DB)
}

type Table struct {
	Game      Properties
	Options   Properties
	Dnd       Properties
	Campaign  Properties
	Weapons   Properties
	Locations Properties
	Load      Properties
	Save      Properties
}

// Creates a file called game.db with GameTable struct
func (T *GameTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

// Updates the GameTable struct
func (T *GameTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable GameTable
	db.Model(&existingGameTable).Update("GameData", game)
}

// Inserts values into struct
func (T *GameTable) Insert(db *gorm.DB) {

}

// Makes queries from the GameTable struct
func (T *GameTable) Query(db *gorm.DB) {

}

// Deletes values from GameTable struct
func (T *GameTable) Delete(db *gorm.DB) {

}

func (T *OptionsTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *OptionsTable) Update(db *gorm.DB, game interface{}) {
	var existingOptionsTable OptionsTable
	db.Model(&existingOptionsTable).Update("OptionsData", game)
}

func (T *OptionsTable) Insert(db *gorm.DB) {

}

func (T *OptionsTable) Query(db *gorm.DB) {

}

func (T *OptionsTable) Delete(db *gorm.DB) {

}

func (T *DndTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *DndTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable DndTable
	db.Model(&existingGameTable).Update("DndData", game)
}

func (T *DndTable) Insert(db *gorm.DB) {

}

func (T *DndTable) Query(db *gorm.DB) {

}

func (T *DndTable) Delete(db *gorm.DB) {

}

func (T *CampaignTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *CampaignTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable CampaignTable
	db.Model(&existingGameTable).Update("CampaignData", game)
}

func (T *CampaignTable) Insert(db *gorm.DB) {

}

func (T *CampaignTable) Query(db *gorm.DB) {

}

func (T *CampaignTable) Delete(db *gorm.DB) {

}

func (T *WeaponsTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *WeaponsTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable WeaponsTable
	db.Model(&existingGameTable).Update("WeaponsData", game)
}

func (T *WeaponsTable) Insert(db *gorm.DB) {

}

func (T *WeaponsTable) Query(db *gorm.DB) {

}

func (T *WeaponsTable) Delete(db *gorm.DB) {

}

func (T *LocationsTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *LocationsTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable LocationsTable
	db.Model(&existingGameTable).Update("LocationData", game)
}

func (T *LocationsTable) Insert(db *gorm.DB) {

}

func (T *LocationsTable) Query(db *gorm.DB) {

}

func (T *LocationsTable) Delete(db *gorm.DB) {

}

//

func (T *LoadTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *LoadTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable LoadTable
	db.Model(&existingGameTable).Update("LoadData", game)
}

func (T *LoadTable) Insert(db *gorm.DB) {

}

func (T *LoadTable) Query(db *gorm.DB) {

}

func (T *LoadTable) Delete(db *gorm.DB) {

}

func (T *SaveTable) Init(db *gorm.DB) {
	db.AutoMigrate(T)
}

func (T *SaveTable) Update(db *gorm.DB, game interface{}) {
	var existingGameTable SaveTable
	db.Model(&existingGameTable).Update("SaveData", game)
}

func (T *SaveTable) Insert(db *gorm.DB) {

}

func (T *SaveTable) Query(db *gorm.DB) {

}

func (T *SaveTable) Delete(db *gorm.DB) {

}

// SerializeGame converts the Game struct into JSON for storage
func Serialize(g interface{}) (string, error) {
	gameJSON, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(gameJSON), nil
}

// DeserializeGame converts JSON back into a Game struct
func Deserialize(data string, out interface{}) error {
	return json.Unmarshal([]byte(data), out)
}
