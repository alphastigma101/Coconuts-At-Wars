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
	Players map[uint][]*Player  // Represents all the players that will be searched by the id field
	Actor   map[string][]string // Three keys and values: Sprite: path-to-it, Death: path-to-it,
}

type enemies struct {
	Health int
	Weapon bool
	Actors map[string]string // Stores the images of monsters that will be either rendered in 2D or 3D
}

// Struct that keeps track of where the user is at in the game and will load them there
type gameState struct {
	Dungeon      bool
	IceLand      bool
	HotLand      bool
	TropicalLand bool
}

type Dnd struct {
	players   Player
	Weapons   map[string]string
	NPCs      map[string]string
	Monsters  enemies
	GameState gameState
}

// During the dnd, if the user clicks a certain button, it will pop up a menu
// They can either adjust game volume, exit, continue, and save and connect
type gameOptions interface {
	Options() interface{}
}

func Campaign(d *Dnd) {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Locations() {
	panic("Function has not been implemented yet!")
}
