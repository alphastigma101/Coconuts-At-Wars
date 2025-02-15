package main

import (
	"time"
	// My own modules
	"github.com/alphastigma101/Coconuts-At-Wars/input_handler"
	// 3D models
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
)

// GameMode represents the dimension type (2D or 3D)
type GameMode int

const (
	Mode2D GameMode = iota
	Mode3D
)

// GameOptions holds configuration for the game
type GameOptions struct {
	Mode GameMode
}

// Renderer interface defines what any rendering system must implement
type Renderer interface {
	Initialize() error
	Render() error
	Cleanup() error
}

// Game2D implements 2D rendering
type Game2D struct {
	options *GameOptions
}

// Cleanup implements Renderer.
func (g *Game2D) Cleanup() error {
	panic("unimplemented")
}

// Render implements Renderer.
func (g *Game2D) Render() error {
	panic("unimplemented")
}

// Game3D implements 3D rendering
type Game3D struct {
	options *GameOptions
}

// Cleanup implements Renderer.
func (g *Game3D) Cleanup() error {
	panic("unimplemented")
}

// Render implements Renderer.
func (g *Game3D) Render() error {
	panic("unimplemented")
}

// TitleScreen manages the main menu
type TitleScreen struct {
	renderer Renderer
	options  *GameOptions
}

// Initialize the 2D renderer
func (g *Game2D) Initialize() error {
	// Initialize 2D graphics (could use Ebiten here)
	return nil
}

// Initialize the 3D renderer
func (g *Game3D) Initialize() error {
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
func CreateRenderer(options *GameOptions) Renderer {
	switch options.Mode {
	case Mode3D:
		return &Game3D{options: options}
	default:
		return &Game2D{options: options}
	}
}

func main() {
	actor := input_handler.GameActor{
		Health: 100,
	}
	// Need a handful of if statements that pull in the start menu of the game
	// It will display the game with three coconuts huddle around
	// Two of them will have army helments while the other coconut on the far right will have
	// A mullet. Everytime the user starts up the game the coconut with the mullet will have different cosmetics
	// Such as googles, a cone on its head, etc
	// They will be huddled around a gray-ish steel circlular object which the landscape will be mud while airplanes fly around
	// By default it will be rendered as 2D graphics
	// Similar to inheritance because an interface needs some object that implements
	// it's properties
	handler := &input_handler.InputHandler{
		ButtonX: &input_handler.JumpCommand{},
		ButtonY: &input_handler.FireCommand{},
		ButtonA: &input_handler.DuckCommand{},
		ButtonB: &input_handler.ReloadCommand{},
	}

	command := input_handler.HandleInput(handler)
	if command != nil {
		command.Execute(&actor)
	}
}
