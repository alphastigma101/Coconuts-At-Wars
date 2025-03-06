package layout

import (
	"math"
	"os"

	r1 "github.com/gen2brain/raylib-go/raylib"
)

// Used by the database
type Properties interface {
	Init(table interface{}, Game interface{}) (interface{}, interface{})
	Update(game interface{})
}

// Use this function for collision in 2D:
//
//	CheckCollisionLines
//
// CheckCollisionPointLine
//
//	CheckCollisionPointPoly
type VertexBuffer *r1.VertexBuffer
type DPI *r1.Vector2 // Dots Per Inch. Has to do with using a mouse
type Texture2D *r1.Texture2D
type Sound *r1.Sound
type Shader *r1.Shader
type Render2D *r1.RenderTexture2D
type Music *r1.Music
type ModelAnimation *r1.ModelAnimation
type Material *r1.Material
type Image *r1.Image
type Position *r1.Vector2
type Scene *r1.DrawCall

// Store Audio data
// Raylib Functions:
// AttachAudioMixedProcessor
type Audio *r1.AudioBuffer
type AudioProcessor *r1.AudioProcessor
type AudioStream *r1.AudioStream

// Camera
// Use this function to rotate the camera
// CameraRoll
// Use this to move the camera up:
// CameraMoveUp
type Camera2D *r1.Camera2D // A Camera Struct meant for 2D. Use it with: BeginMode2D function
type Camera3D *r1.Camera3D

// Events
type Event *r1.AutomationEvent
type EventList *r1.AutomationEventList

// Drawing Stuff Ontop of Images and or Creating Sprites
// Colors
var Black = r1.Black
var White = r1.White
var Red = r1.Red
var Gray = r1.Gray

type Rectangle r1.Rectangle
type Color r1.Color

func BeginDrawing() {
	r1.BeginDrawing()
}
func EndDrawing() {
	r1.EndDrawing()
}
func ClearBackground(col Color) {
	// Convert your custom Color type to r1.Color
	r1Color := r1.Color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}
	// Pass the converted color to r1.ClearBackground
	r1.ClearBackground(r1Color)
}
func ColorAlpha(col Color, alpha float32) Color {
	r1Color := r1.Color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}
	rgba := r1.ColorAlpha(r1Color, alpha)
	return Color{R: rgba.R, G: rgba.G, B: rgba.B, A: rgba.A}
}

// Images and Textures
func LoadImage(file string) *Image {
	r1Image := r1.LoadImage(file)
	img := (Image)(r1Image)
	return &img
}
func UnloadTexture(texture Texture2D) {
	r1.UnloadTexture(*texture)
}
func UnloadImage(img Image) {
	r1img := r1.Image(*img)
	r1.UnloadImage(&r1img)
}
func LoadTextureFromImage(img Image) Texture2D {
	r1Img := r1.Image(*img)
	r1Texture := r1.LoadTextureFromImage(&r1Img)
	return Texture2D(&r1Texture)
}
func DrawTexture(texture Texture2D, posX int32, posY int32, col Color) {
	r1Texture := r1.Texture2D(*texture)
	r1Color := r1.Color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}
	r1.DrawTexture(r1Texture, posX, posY, r1Color)
}
func DrawText(text string, posX int32, posY int32, fontSize int32, col Color) {
	r1Color := r1.Color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}
	r1.DrawText(text, posX, posY, fontSize, r1Color)
}
func DrawRectangle(posX int32, posY int32, width int32, height int32, col Color) {
	r1Color := r1.Color{
		R: uint8(col.R),
		G: uint8(col.G),
		B: uint8(col.B),
		A: uint8(col.A),
	}
	r1.DrawRectangle(posX, posY, width, height, r1Color)
}

// Key Variables
var KeyDown int = r1.KeyDown
var KeyUp int = r1.KeyUp
var KeyEnter int = r1.KeyEnter

// Key Event Functions
func IsKeyPressed(keyDown int) bool {
	return r1.IsKeyPressed(int32(keyDown))
}

// Getter Methods
// Windows
func WindowShouldClose() bool {
	return r1.WindowShouldClose()
}
func InitWindow(width int32, height int32, title string) {
	r1.InitWindow(width, height, title)
}
func GetScreenHeight() int {
	return r1.GetScreenHeight()
}
func GetScreenWidth() int {
	return r1.GetScreenWidth()
}
func GetWindowPosition() Position {
	r1Pos := r1.GetWindowPosition()
	return &r1Pos
}
func GetWindowScaleDPI() Position {
	r1Pos := r1.GetWindowScaleDPI()
	return &r1Pos
}
func GetTime() float64 {
	return r1.GetTime()
}

// FPS settings
func SetTargetFPS(fps int32) {
	r1.SetTargetFPS(fps)
}

// Audio free-functions
func InitAudioDevice() {
	r1.InitAudioDevice()
}

// 3D Stuff
type Mesh *r1.Mesh // It is apart of the 3D model
// Would want to use this function with it:
// CheckCollisionBoxes
type Box *r1.BoundingBox // Can be used for hitbox that can be placed on the player

type GameRenderer interface {
	InitializeOptionsScreen() interface{}
	//InitializeLoadScreen() interface{}
	InitializeCampaignScreen()
	//InitializeDndScreen() interface{}
}

// An interface that defines a 2D rendering system must implement
type Render interface {
	InitializeTitleScreen() interface{}
	InitializeMainMenuScreen(Game GameRenderer)
}
type mainMenu struct {
	Rectangle Rectangle
	Color     Color
	Active    bool
}

type layout struct {
	MainMenu *mainMenu
}

type Game2D struct {
	VertexBuffer VertexBuffer
	DPI          DPI
	Texture      Texture2D
	// Use LoadSound(fileName string) to load in music
	Sound          Sound
	Shader         Shader
	Render         Render2D // Renders the texture into 2D
	Music          Music
	ModelAnimation ModelAnimation
	Material       Material
	Image          Image
	Scene          Scene
	Audio          Audio          // Create a audio buffer
	AudioProcessor AudioProcessor // Process the audio
	Camera         Camera2D       // A copy of the Camera2D struct
	// Use LoadAudioStream if you're using custom audio splices
	AudioStream AudioStream // Create custom audio streams not bound to a specific file
	// Use LoadAutomationEventList to load in
	Event      Event
	EventList  EventList
	Scenes     map[string]*Game2D // Store all the images that have been scaled
	GameLayout *layout
}

func (g Game2D) InitializeTitleScreen() interface{} {

	// Load and check image
	image := LoadImage("./assests/titlescreen.png")
	if image == nil {
		panic("failed to load titlescreen image")
	}
	texture := LoadTextureFromImage(*image)
	texture.Height = int32(GetScreenHeight())
	texture.Width = int32(GetScreenWidth())
	UnloadImage(*image)
	inputDetected := false
	for !inputDetected {
		if WindowShouldClose() {
			r1.CloseWindow()
			os.Exit(0)
		}

		// Start rendering
		BeginDrawing()
		ClearBackground(Color{R: White.R, G: White.G, B: White.B, A: White.A})

		// Draw title screen and draw the texture after loading it
		DrawTexture(texture, 0, 0, Color{R: White.R, G: White.G, B: White.B, A: White.A})
		DrawText("PRESS ANY BUTTON TO START", 240, 400, 20, Color{R: White.R, G: White.G, B: White.B, A: White.A})
		pressed := r1.GetKeyPressed()
		if pressed != 0 {
			inputDetected = true
		}
		EndDrawing() // Finish drawing
	}
	// Unload texture after input is detected
	UnloadTexture(texture)
	return &g
}

func roundOdd(num int) int {
	if num%25 != 0 {
		num = num + 1
		return roundOdd(num)
	}
	return num
}

func freeResources(images []*Image, textures []Texture2D) {
	/*for i := range images {
		if images[i] != nil {
			UnloadImage(*images[i])
		}
		//if textures[i].ID != 0 {
		//UnloadTexture(textures[i])
		//}
	}*/

}
func (g Game2D) InitializeMainMenuScreen(Game GameRenderer) {
	BeginDrawing()
	imageFiles := []string{
		"./assests/Dry WarZone.jpg",
		"./assests/Frozen WarZone.png",
		"./assests/Muddy WarZone.png",
	}

	// Create array of image pointers
	images := make([]*Image, len(imageFiles))

	// Load all images
	for i, file := range imageFiles {
		images[i] = LoadImage(file)
		if images[i] == nil {
			panic("failed to load image: " + file)
		}
	}

	// Create array for textures
	textures := make([]Texture2D, len(images))

	// Track current image index
	currentImageIndex := 0
	time := 0.0
	intervalImageChange := 0

	// Menu options
	menuOptions := []string{"Campaign", "Load", "Dnd", "Options", "Exit"}
	selectedOption := 0
	exitMenu := false
	init := false
	for !exitMenu && !WindowShouldClose() {
		pressed := r1.GetKeyPressed()
		// Handle events and input
		if r1.IsKeyPressed(int32(KeyDown)) || pressed == 264 {
			selectedOption = (selectedOption + 1) % len(menuOptions)
		} else if IsKeyPressed(KeyUp) || pressed == 265 {
			selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
		} else if IsKeyPressed(KeyEnter) || pressed == 257 {
			// Process selection
			switch menuOptions[selectedOption] {
			case "Campaign":
				freeResources(images, textures)
				Game.InitializeCampaignScreen()
				exitMenu = true
			case "Load":
				freeResources(images, textures)
				//g.Event = EventLoad
				exitMenu = true
			case "Dnd":
				// Call in the Dnd screen function and before loading anyting
				// We need to check and see if DndMode is enabled
				// If not return it back to here and make a warning pop up box
				freeResources(images, textures)
				//g.Event = EventDnd
				exitMenu = true
			case "Options":
				freeResources(images, textures)
				//g.Event = EventOptions
				exitMenu = true
			case "Exit":
				freeResources(images, textures)
				g.InitializeTitleScreen()
			}
		}
		if !init && intervalImageChange == 0 {
			// Initialize the textures
			textures[currentImageIndex] = LoadTextureFromImage(*images[currentImageIndex])
			textures[currentImageIndex].Height = int32(GetScreenHeight())
			textures[currentImageIndex].Width = int32(GetScreenWidth())
			init = true
		}
		if intervalImageChange == 1500 {
			if textures[currentImageIndex].ID != 0 {
				UnloadTexture(textures[currentImageIndex])
			}
			// Move to next image
			currentImageIndex = (currentImageIndex + 1) % len(images)
			// Convert current image to texture
			textures[currentImageIndex] = LoadTextureFromImage(*images[currentImageIndex])
			textures[currentImageIndex].Height = int32(GetScreenHeight())
			textures[currentImageIndex].Width = int32(GetScreenWidth())
			intervalImageChange = 0
			time = GetTime()
			time = float64(roundOdd(int(time)))
			if int(math.Mod(time, 25)) == 0 && time > 0 {
				// TODO: Load in music here for each image
				print("Being executed more than once!\n")
			}
		}
		ClearBackground(Color{R: Black.R, G: Black.G, B: Black.B, A: Black.A})
		DrawTexture(textures[currentImageIndex], 0, 0, Color{R: White.R, G: White.G, B: White.B, A: White.A})

		// Draw menu box
		menuX := int32(GetScreenWidth())/2 - 100
		menuY := int32(GetScreenHeight()) / 2

		// Store menu
		if g.GameLayout == nil {
			g.GameLayout = &layout{}
			g.GameLayout.MainMenu = &mainMenu{
				Rectangle: Rectangle{
					X:      float32(menuX),
					Y:      float32(menuY),
					Width:  200,
					Height: float32(len(menuOptions) * 40),
				},
			}
		}

		// Draw the menu box
		col := Color{R: Black.R, G: Black.G, B: Black.B, A: Black.A}
		DrawRectangle(menuX, menuY, 200, int32(len(menuOptions)*40), ColorAlpha(col, 0.7))

		// Draw menu options
		for i, option := range menuOptions {
			textColor := Color{R: White.R, G: White.G, B: White.B, A: White.A}
			if i == selectedOption {
				textColor = Color{R: r1.Red.R, G: r1.Red.G, B: r1.Red.B, A: r1.Red.A}
				col = Color{R: Gray.R, G: Gray.G, B: Gray.B, A: Gray.A}
				DrawRectangle(menuX, menuY+int32(i*40), 200, 40, ColorAlpha(col, 0.3))
			}
			DrawText(option, menuX+20, menuY+int32(i*40)+10, 20, textColor)
		}
		intervalImageChange = intervalImageChange + 1
		EndDrawing()
	}
}

func GetGame2D() Game2D {
	return Game2D{}
}

// Game3D implements 3D rendering
type game3D struct {
	//Transform      r1.Transform
	Music          Music
	ModelAnimation ModelAnimation
	Mesh           Mesh
	Material       Material
	Image          Image
	Scene          Scene
	Gui            Box
	Audio          Audio          // Create a audio buffer
	AudioProcessor AudioProcessor // Process the audio
	Camera         Camera3D       // A copy of the Camera3D struct
	// Use LoadAudioStream if you're using custom audio splices
	AudioStream r1.AudioStream // Create custom audio streams not bound to a specific file
	// Use LoadAutomationEventList to load in
	Event     Event
	EventList EventList
	Scenes    map[string]*game3D
}
type Game3D game3D

func (g game3D) InitializeTitleScreen() interface{} {
	return nil
}

func (g game3D) InitializeMainMenuScreen(game interface{}) interface{} {
	return nil
}

func GetGame3D() game3D {
	return game3D{}
}
