package options

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	// 2D modules
	//"github.com/hajimehoshi/ebiten/v2"
	//"bytes"
	//"image"
	//"log"
)

// GameMode represents the dimension type (2D or 3D)
type gameMode int

// GameOptions holds configuration for the game
type gameOptions struct {
	Mode gameMode
}

// Game2D implements 2D rendering
type game2D struct {
	//a := app.App()
	//scene := core.NewNode()
}

// Game3D implements 3D rendering
type game3D struct {
	app      *app.Application
	scene    *core.Node
	camera   *camera.Camera
	renderer *renderer.Renderer
	//window    *window.Window
	//gui       *gui.GUI
	lightNode *core.Node
}

// TitleScreen manages the main menu
type titleScreen struct {
	renderer Renderer
	options  gameOptions
}

type Game struct {
	Mode    gameMode
	Options gameOptions
	Game2D  game2D
	Game3D  game3D
	Title   titleScreen
}

// Renderer interface defines what any rendering system must implement
type Renderer interface {
	Initialize() error
	Render() error
	Cleanup() error
}

// Cleanup implements Renderer for 2D.
func (g *Game) Cleanup() error {
	panic("unimplemented")
}

// Cleanup implements Renderer.
func (g *game2D) Cleanup() error {
	panic("unimplemented")
}

// Render implements Renderer.
func (g *game2D) Render() error {
	panic("unimplemented")
}

// Cleanup implements Renderer.
func (g *game3D) Cleanup() error {
	panic("unimplemented")
}

// Render implements Renderer.
func (g *game3D) Render() error {
	panic("unimplemented")
}

// Render implements Renderer for 2D.
func (g *Game) Render() error {
	panic("unimplemented")
}

// Initialize the 2D renderer
func (g *game2D) Initialize() error {
	// Initialize 2D graphics (could use Ebiten here)
	return nil
}

// Initialize the 3D renderer
func (g *game3D) Initialize() error {
	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Create a blue torus and add it to the scene
	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(geom, mat)
	scene.Add(mesh)

	// Create and add a button to the scene
	btn := gui.NewButton("Make Red")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		mat.SetColor(math32.NewColor("DarkRed"))
	})
	scene.Add(btn)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
	return nil
}

// CreateRenderer factory function to create appropriate renderer
func CreateGame(options *Game) *Game {
	switch options.Mode {
	case 1:
		init := &game3D{
			app:    app.App(),
			scene:  core.NewNode(),
			camera: camera.New(1),
			//renderer: renderer.New(),
			lightNode: core.NewNode(),
		}
		options.Game3D = *init
		return options
	default:
		//return &game2D{}
	}
	panic("Failed to initalize the game in 2D or 3D!")
}
