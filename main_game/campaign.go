package campaign

type Campaign interface {
	Locations() // represents the locations
	Weapons()   // represents the weapons
	World()     // represents the world which will be passing the game2D to it
}
