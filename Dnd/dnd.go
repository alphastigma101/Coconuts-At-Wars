package dnd

// Simple bag that the player can access
type bag struct {
	healing []string
	weapons []string
	armor   []string
}

type Players struct {
	Bag bag // A struct that Represents the current player's bag
	id  int // A unique id that gets assigned to each player
}

type Dnd struct {
	players Players
	Images  []string // Stores the images that will be either rendered in 2D or 3D
}

func (d *Dnd) Campaign() {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *Dnd) Locations() {
	panic("Function has not been implemented yet!")
}
