package render2d

import (
	// Printing
	"fmt"
	// 2D modules
	"github.com/g3n/engine/math32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rakyll/statik/fs"

	//"image"
	"image/png"
)

// Game2D implements 2D rendering
type Game2D struct {
	// Going to need to execute RunGame and RunGameWithOptions but do not execute it twice
	// This is the function that should be called: RunGameWithOptions
	Game         *ebiten.Game
	Options      *ebiten.RunGameOptions // Options that are used with the game
	screen       *ebiten.Image          // Store the settings of magnifying the image
	screenWidth  int                    // Set the width
	screenHeight int                    // Set the height of the screen
	scale        float32                // Scale the image
	Scenes       map[string]*Game2D     // Store all the images that have been scaled
}

func Init2DStruct() *Game2D {
	return &Game2D{
		screenWidth:  0,
		screenHeight: 0,
	}
}

// Initialize the 2D renderer
func (g Game2D) InitializeTitleScreen() error {
	// Set up the screen dimensions
	screenWidth := 800
	screenHeight := 600
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Coconuts At Wars")

	// Load embedded resources
	statikFS, err := fs.New()
	if err != nil {
		return fmt.Errorf("failed to create statik fs: %v", err)
	}

	// Open the embedded file
	r, err := statikFS.Open("assests/titlescreen.png")
	if err != nil {
		return fmt.Errorf("failed to open embedded title screen: %v", err)
	}
	defer r.Close()

	// Decode the image
	img, err := png.Decode(r)
	if err != nil {
		return fmt.Errorf("failed to decode title screen: %v", err)
	}

	// Convert to ebiten.Image
	titleImg := ebiten.NewImageFromImage(img)

	// Create screen buffer
	g.screen = ebiten.NewImage(screenWidth, screenHeight)

	// Scale image to fit screen
	bounds := titleImg.Bounds()
	scaleX := float32(screenWidth) / float32(bounds.Dx())
	scaleY := float32(screenHeight) / float32(bounds.Dy())

	// Store scaled image and dimensions
	g.screen = titleImg
	g.scale = math32.Min(scaleX, scaleY)
	g.screenWidth = screenWidth
	g.screenHeight = screenHeight
	g.Scenes["title"] = &g
	return nil
}

func (g Game2D) Update() error {
	return nil
}

func (g Game2D) Draw(screen *ebiten.Image) {
	// If we have a screen buffer, draw it
	if g.screen != nil {
		op := &ebiten.DrawImageOptions{}
		// Draw the buffer to the screen
		screen.DrawImage(g.screen, op)
	}
	// Update our screen reference
	g.screen = screen
}

func (g Game2D) Render() error {
	// Call in the gui within the if statement
	// Rendering will be getting the images from campaign/Locations
	// And adding it as the player progresses through the game
	// Needs to be scaled to 2D
	panic("unimplemented")
}

func (g Game2D) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//render, ok := options.GetGame2D(g.Options)
	return 800, 600 // example dimensions
}

func (g Game2D) CleanUp() error {
	// Call in the gui within the if statement
	// Rendering will be getting the images from campaign/Locations
	// And adding it as the player progresses through the game
	// Needs to be scaled to 2D
	panic("unimplemented")
}
