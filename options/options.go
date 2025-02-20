package options

import (
	"fmt"

	"github.com/alphastigma101/Coconuts-At-Wars/render2d"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/alphastigma101/Coconuts-At-Wars/render3d"
)

// GameMode represents the dimension type (2D or 3D)
type gameMode int
type dndMode int

// An interface that defines a 2D rendering system must implement
type render2D interface {
	InitializeTitleScreen() error
	Update() error
	Draw(screen *ebiten.Image) // Will be a pointer that points to the Image struct in ebiten
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Render() error
	CleanUp() error
}

// An interface that defines a 3D rendering system must implement
type render3D interface {
	InitializeTitleScreen() error
	//Update() error
	//Draw(screen *ebiten.Image) // Will be a pointer that points to the Image struct in ebiten
	//Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type Options struct {
	GameMode gameMode
	DndMode  dndMode
	Game2D   render2D
	Game3D   render3D
}

func GetGame2D(o *Options) (render2d.Game2D, error) {
	if o == nil {
		return render2d.Game2D{}, fmt.Errorf("options is nil")
	}

	// Use type assertion to check if o.Game2D implements game2D
	g2d, ok := o.Game2D.(render2d.Game2D)
	if !ok {
		return render2d.Game2D{}, fmt.Errorf("Game2D does not implement game2D")
	}
	return g2d, nil
}

// CreateRenderer factory function to create appropriate renderer
func UpdateOptions(game *Options) *Options {
	if game.GameMode == 1 {
		//game.Game3D = *render3d.Init3DStruct()
		return game
	} else {
		game.Game2D = *render2d.Init2DStruct()
		return game
	}
}
