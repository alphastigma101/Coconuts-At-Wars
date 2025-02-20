package render3d

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
)

// Game3D implements 3D rendering
type game3D struct {
	App      *app.Application
	Scene    *core.Node
	Camera   *camera.Camera
	Renderer *renderer.Renderer
	//window    *window.Window
	gui       *gui.Window
	LightNode *core.Node
	Geom      geometry.Geometry
	Mat       material.Standard
	//mesh := graphic.NewMesh(geom, mat)
	Scenes map[string]*game3D
}

// Initialize the 3D renderer
func (g game3D) InitializeTitleScreen() error {
	// Set the scene to be managed by the gui manager
	gui.Manager().Set(g.Scene)

	// Create perspective camera
	cam := g.Camera
	cam.SetPosition(0, 0, 3)
	g.Scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := g.App.GetSize()
		g.App.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	g.App.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Create a blue torus and add it to the scene
	// TODO: g.Geom and g.Mat and mesh create an object and will need to be added to the scene
	// This object must be stored inside of a map
	g.Geom = *geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	g.Mat = *material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(&g.Geom, &g.Mat)
	g.Scene.Add(mesh)

	// Create and add a button to the scene
	btn := gui.NewButton("Make Red")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		g.Mat.SetColor(math32.NewColor("DarkRed"))
	})
	g.Scene.Add(btn)

	// Create and add lights to the scene
	g.Scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	g.Scene.Add(pointLight)

	// Create and add an axis helper to the scene
	g.Scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	g.App.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	g.App.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		g.App.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(g.Scene, cam)
	})
	return nil
}

func Init3DStruct() *game3D {
	return &game3D{
		App:    app.App(),
		Scene:  core.NewNode(),
		Camera: camera.New(1),
		//renderer: renderer.New(),
		LightNode: core.NewNode(),
	}
}
