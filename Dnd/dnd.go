package dnd

// Simple bag that the player can access
type bag struct {
	Healing  []string
	Weapons  []string
	Armor    []string
	Location map[string]string
}

type Player struct {
	Bag    bag          // A struct that Represents the current player's bag
	id     uint         // Initialize the keys from the database to represent the player. Generate the id for npcs
	coords map[uint]int // represents the coordinates where the player is on the map
	// There is no database table for players, so store it inside the struct
	Players map[uint][]*Player // Represents all the players that will be searched by the id field
	Actor   interface{}        // There is a struct called gameActor or similar to that, it will be initialized with this field
}

type Dnd struct {
	players  Player
	Monsters map[string]string // Stores the images that will be either rendered in 2D or 3D
	Weapons  map[string]string
	NPCs     map[string]string
}

// Game struct has a Dnd field and there is a dnd table, therefore we can make this a type instead
// Of storing it inside the struct
type SavedData map[uint][]*Dnd // Represents the progress through the dnd game

func (d *Dnd) Campaign() {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Locations() {
	panic("Function has not been implemented yet!")
}
