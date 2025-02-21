package options

import (
	"fmt"

	r1 "github.com/gen2brain/raylib-go/raylib"
)

// GameMode represents the dimension type (2D or 3D)
type gameMode int
type dndMode int

// An interface that defines a 2D rendering system must implement
type render interface {
	InitializeTitleScreen() interface{}
}

type game2D struct {
	VertexBuffer *r1.VertexBuffer
	Position     *r1.Vector2
	DPI          *r1.Vector2 // Store the window dpi factor
	Texture      *r1.Texture2D
	// Use LoadSound(fileName string) to load in music
	Sound          *r1.Sound
	Shader         *r1.Shader
	Render         *r1.RenderTexture2D
	Music          *r1.Music
	ModelAnimation *r1.ModelAnimation
	Mesh           *r1.Mesh
	Material       *r1.Material
	Image          *r1.Image
	Scene          *r1.DrawCall
	Gui            *r1.BoundingBox
	Audio          *r1.AudioBuffer    // Create a audio buffer
	AudioProcessor *r1.AudioProcessor // Process the audio
	Camera         *r1.Camera2D       // A copy of the Camera2D struct
	// Use LoadAudioStream if you're using custom audio splices
	AudioStream *r1.AudioStream // Create custom audio streams not bound to a specific file
	// Use LoadAutomationEventList to load in
	Event     *r1.AutomationEvent
	EventList *r1.AutomationEventList
	Scenes    map[string]*game2D // Store all the images that have been scaled
}

// Game3D implements 3D rendering
type game3D struct {
	Transform      *r1.Transform
	Music          *r1.Music
	ModelAnimation *r1.ModelAnimation
	Mesh           *r1.Mesh
	Material       *r1.Material
	Image          *r1.Image
	Scene          *r1.DrawCall
	Gui            *r1.BoundingBox
	Audio          *r1.AudioBuffer    // Create a audio buffer
	AudioProcessor *r1.AudioProcessor // Process the audio
	Camera         *r1.Camera3D       // A copy of the Camera3D struct
	// Use LoadAudioStream if you're using custom audio splices
	AudioStream *r1.AudioStream // Create custom audio streams not bound to a specific file
	// Use LoadAutomationEventList to load in
	Event     *r1.AutomationEvent
	EventList *r1.AutomationEventList
	Scenes    map[string]*game3D
}

type Options struct {
	GameMode gameMode
	DndMode  dndMode
	Game2D   render
	Game3D   render
}

func (g game2D) InitializeTitleScreen() interface{} {
	// Initialize window first
	r1.InitWindow(800, 450, "Coconuts At Wars")
	r1.SetTargetFPS(60)

	// Get window info
	dpi := r1.GetWindowScaleDPI()
	pos := r1.GetWindowPosition()

	// Load and check image
	image := r1.LoadImage("./assets/titlescreen.png")
	if image == nil {
		panic("failed to load titlescreen image")
	}

	// Convert image to texture for rendering
	texture := r1.LoadTextureFromImage(image)

	// Free the image data since we've converted it to texture
	r1.UnloadImage(image)

	// Store all the data in the game2D struct
	g.Position = &pos
	g.DPI = &dpi
	g.Texture = &texture
	g.Image = image

	// Initialize scenes map if it doesn't exist
	if g.Scenes == nil {
		g.Scenes = make(map[string]*game2D)
	}
	g.Scenes["title"] = &g
	return &g
}

func GetGame2D(o *Options) (game2D, error) {
	if o == nil {
		return game2D{}, fmt.Errorf("options is nil")
	}

	// Use type assertion to check if o.Game2D implements game2D
	g2d, ok := o.Game2D.(game2D)
	if !ok {
		return game2D{}, fmt.Errorf("Game2D does not implement game2D")
	}
	return g2d, nil
}
func (g game3D) InitializeTitleScreen() interface{} {
	return nil
}
func GetGame3D(o *Options) (game3D, error) {
	if o == nil {
		return game3D{}, fmt.Errorf("options is nil")
	}

	// Use type assertion to check if o.Game2D implements game2D
	g3d, ok := o.Game3D.(game3D)
	if !ok {
		return game3D{}, fmt.Errorf("Game2D does not implement game2D")
	}
	return g3d, nil
}

// CreateRenderer factory function to create appropriate renderer
func UpdateOptions(game *Options) *Options {
	if game.GameMode == 1 {
		// Need to update the Options Table
		game.Game3D = &game3D{}
		return game
	} else {
		game.Game2D = &game2D{}
		return game
	}
}
