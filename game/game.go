/*
This is the game module
*/
package game

// This Module holds the Game struct which consists of nested structs that keeps track of the user's data
import (

	// My modules
	"fmt"

	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	coop "github.com/alphastigma101/Coconuts-At-Wars/cooperative"
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
	"github.com/alphastigma101/Coconuts-At-Wars/vehicles"
)

// TODO: Need to remove the Load Data table. There is no need for it when the GameTable already does that
// Will be using gameState struct which has a series of boolean fields that will be used as a pivot
// It will allow the user to load in certain places once they are unlocked

type DataBaseProperties Layout.Properties

type Vehicles struct {
	Plane vehicles.Plane
	Boat  vehicles.Boat
	Car   vehicles.Car
	Truck vehicles.Truck
}

type enemies struct {
	Health int
	Weapon bool
}

type Campaign struct {
	Players dnd.Player
	Enemies map[string]enemies // Stores the images that will be either rendered in 2D or 3D
}

// Struct that keeps track of where the user is at in the game and will load them there
type gameState struct {
	DndTutorial     bool
	RegularTutorial bool
	DdayArea        bool
	WasteLandArea   bool
	RiverArea       bool
}
type GameState gameState

// Struct that will keep track of the damage and the position of the player
type gameActor struct {
	Health   int
	Position Layout.Position
}

type Actor gameActor

// A struct that holds a series of other stand alone structs forming some kind of web
type Game struct {
	Database    Layout.Table      // A struct that can be used for easy access to the database table
	Options     *options.Options  // A struct that can enable coop, 2D or 3D gaming
	Dnd         *dnd.Dnd          // A struct that has the Dnd related fields to function on its own
	Cooperative *coop.Cooperative // A struct that has Coop related fields to function on its own
	Players     *dnd.Player       // A struct that is used to keep track of the health, weapons, and location
	GameActor   *Actor            // This struct won't be needed. Player field will be used instead
	Game2D      Layout.Render
	Game3D      Layout.Render
	GameState   GameState // A struct that keeps track of where the player is at and can load the player at that location
	Vehicles    Vehicles  // A struct that has all the vehicles aliases that the user can use
}

// During the campaign, if the user clicks a certain button, it will pop up a menu
// They can either adjust game volume, exit, continue, and connect
type GameOptions interface {
	Options() interface{}
}

// Getter Functions

func (g *Game) IsDndEnabled() bool {
	return g.Options.DndMode == 1
}

func (g *Game) IsCoopEnabled() bool {
	return g.Options.Coop == 1
}

// This function will render in the D-Day map area
// The sprites will need to be loaded in and so does the collision
func (g *Game) dDay() {

}

func (g *Game) River() {

}

func (g *Game) RegularTutorial() {
	// Need to check and see what the user has enabled in their settings to determine if 2D or 3D is enabled
	// We also need to check constantly to see if the user has pressed a certain key that will open the game options
	// Map needs to be loaded in called forest
	// It is made by krita, and it has animation layers in it
	// The animation layers has different coconuts that will need to have different dialogue above their head whenever the
	// player is near them.
	// Collision also needs to be implemented into the map so the user's sprite does not go out of bounds
	// Will be accessing the field type called Player
	// The GameTable key needs to be assigned with the player id struct because it is unique
	// Need to load in the sprite called baby coconut
	// Need to get the positions for it which will be vector 2
}

// Free function that allows the user to choose a place to load at if they unlocked it
// This will be determined by the GameStateStruct
func (d *Game) UserContinue() {
	if d.GameState.RiverArea {

	} else if d.GameState.WasteLandArea {

	} else if d.GameState.DdayArea {

	}
}

func (d *Game) Wepaons() {
	panic("Function has not been implemented yet!")
}

func (d *Game) Locations() {
	panic("Function has not been implemented yet!")
}

// Functions that initalize the main menu screen

func (g *Game) InitializeCampaignScreen() {
	Layout.BeginDrawing()
	menuOptions := []string{"Continue", "New Game", "Tutorial", "Exit"}
	selectedOption := 0
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		for !Layout.WindowShouldClose() {
			Layout.BeginDrawing() // Start drawing at the beginning of each frame
			pressed := Layout.GetKeyPressed()
			if Layout.IsKeyPressed(Layout.KeyDown) || pressed == 264 {
				selectedOption = (selectedOption + 1) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyUp) || pressed == 265 {
				selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyEnter) || pressed == 257 {
				// Process selection
				switch menuOptions[selectedOption] {
				case "Exit":
					Layout.EndDrawing()
					return
				case "Continue":
					if g.GameState.DdayArea {
						Layout.EndDrawing()
						g.UserContinue()
						return
					}
				case "New Game":
					Layout.EndDrawing()
					g.dDay()
					return
				case "Tutorial":
					// TODO: Tutorial has not been implemented yet.
					// It will be implemented later on
					break
				}
			}
			// Count visible options for proper box scaling
			visibleOptions := 0
			for _, option := range menuOptions {
				if option == "Continue" && !g.GameState.DdayArea {
					continue
				}
				visibleOptions++
			}
			// Position menu on the right side as specified
			menuX := int32(Layout.GetScreenWidth())/2 - 100 + (int32(len(menuOptions) * 40))
			menuY := int32(Layout.GetScreenHeight()) / 2
			// Draw the menu box scaled to visible options
			col := Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A}
			Layout.DrawRectangle(menuX, menuY, 200, int32(visibleOptions*40), Layout.ColorAlpha(col, 0.7))
			// Draw menu options with proper positioning
			visibleIndex := 0
			for i, option := range menuOptions {
				if option == "Continue" && !g.GameState.DdayArea {
					continue
				}
				textColor := Layout.Color{R: Layout.White.R, G: Layout.White.G, B: Layout.White.B, A: Layout.White.A}
				if i == selectedOption {
					textColor = Layout.Color{R: Layout.Red.R, G: Layout.Red.G, B: Layout.Red.B, A: Layout.Red.A}
					col = Layout.Color{R: Layout.Gray.R, G: Layout.Gray.G, B: Layout.Gray.B, A: Layout.Gray.A}
					Layout.DrawRectangle(menuX, menuY+int32(visibleIndex*40), 200, 40, Layout.ColorAlpha(col, 0.3))
				}
				Layout.DrawText(option, menuX+20, menuY+int32(visibleIndex*40)+10, 20, textColor)
				visibleIndex++
			}
			Layout.EndDrawing() // End drawing at the end of each frame
		}
	}
}

func (g *Game) InitializeOptionsScreen(table *Layout.Table) (interface{}, interface{}) {
	menuOptions := []string{"Dnd Option", "Game Mode", "Coop Mode", "Sound", "Exit"}
	selectedOption := 0
	// Add a volume value to track sound level (0-100)
	soundVolume := 50 // Default to 50%
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		for !Layout.WindowShouldClose() {
			Layout.BeginDrawing() // Start drawing at the beginning of each frame
			pressed := Layout.GetKeyPressed()
			if Layout.IsKeyPressed(Layout.KeyDown) || pressed == 264 {
				selectedOption = (selectedOption + 1) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyUp) || pressed == 265 {
				selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyEnter) || pressed == 257 {
				switch menuOptions[selectedOption] {
				case "Exit":
					//table.Options.Update(g)
					Layout.EndDrawing() // End drawing before returning
					return g, table
				}
			} else if Layout.IsKeyPressed(Layout.KeyLeft) || pressed == 263 {
				// Decrease value when the Left Arrow Key is pressed
				switch menuOptions[selectedOption] {
				case "Dnd Option":
					g.Options.DndMode = 0
				case "Game Mode":
					g.Options.GameMode = 0
				case "Coop Mode":
					g.Options.Coop = 0
				case "Sound":
					// Decrease volume by 5%, but not below 0
					soundVolume = max(0, soundVolume-5)
					Layout.SetMasterVolume(float32(soundVolume) / 100.0)
				}
			} else if Layout.IsKeyPressed(Layout.KeyRight) || pressed == 262 {
				// Increase value when the Left Arrow Key is pressed
				switch menuOptions[selectedOption] {
				case "Dnd Option":
					g.Options.DndMode = 1
				case "Game Mode":
					g.Options.GameMode = 1
				case "Coop Mode":
					g.Options.Coop = 1
				case "Sound":
					// Increase volume by 5%, but not above 100
					soundVolume = min(100, soundVolume+5)
					Layout.SetMasterVolume(float32(soundVolume) / 100.0)
				}
				//table.Options.Update(g)
			}
			// Get the dimensions for the black box
			menuX := int32(Layout.GetScreenWidth())/2 - 100 + (int32(len(menuOptions) * 40))
			menuY := int32(Layout.GetScreenHeight()) / 2
			// Draw the menu box
			col := Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A}
			Layout.DrawRectangle(menuX, menuY, 200, int32(len(menuOptions)*40), Layout.ColorAlpha(col, 0.7))
			// Draw menu options
			for i, option := range menuOptions {
				textColor := Layout.Color{R: Layout.White.R, G: Layout.White.G, B: Layout.White.B, A: Layout.White.A}
				if i == selectedOption {
					textColor = Layout.Color{R: Layout.Red.R, G: Layout.Red.G, B: Layout.Red.B, A: Layout.Red.A}
					col = Layout.Color{R: Layout.Gray.R, G: Layout.Gray.G, B: Layout.Gray.B, A: Layout.Gray.A}
					Layout.DrawRectangle(menuX, menuY+int32(i*40), 200, 40, Layout.ColorAlpha(col, 0.3))
				}
				// Draw option text with toggle indicators
				switch option {
				case "Dnd Option":
					dndStatus := "Off"
					if g.Options.DndMode == 1 {
						dndStatus = "On"
					}
					Layout.DrawText(option+": <"+dndStatus+">", menuX+20, menuY+int32(i*40)+10, 20, textColor)
				case "Game Mode":
					modeStatus := "2D"
					if g.Options.GameMode == 1 {
						modeStatus = "3D"
					}
					Layout.DrawText(option+": <"+modeStatus+">", menuX+20, menuY+int32(i*40)+10, 20, textColor)
				case "Coop Mode":
					modeStatus := "Off"
					if g.Options.Coop == 1 {
						modeStatus = "On"
					}
					Layout.DrawText(option+": <"+modeStatus+">", menuX+20, menuY+int32(i*40)+10, 20, textColor)
				case "Sound":
					// Draw volume bar
					Layout.DrawText(option+": ", menuX+20, menuY+int32(i*40)+10, 20, textColor)
					// Calculate bar position and width
					barX := menuX + 90
					barY := menuY + int32(i*40) + 18
					fullBarWidth := 80
					// Draw empty bar background
					Layout.DrawRectangle(barX, barY, int32(fullBarWidth), 6, Layout.ColorAlpha(Layout.Color{R: Layout.Gray.R, G: Layout.Gray.G, B: Layout.Gray.B, A: Layout.Gray.A}, 0.5))
					// Draw filled portion based on current volume
					filledWidth := int32(float32(fullBarWidth) * float32(soundVolume) / 100.0)
					Layout.DrawRectangle(barX, barY, filledWidth, 6, textColor)
					// Draw volume percentage text
					volText := fmt.Sprintf("%d%%", soundVolume)
					Layout.DrawText(volText, barX+int32(fullBarWidth)+5, barY-3, 16, textColor)
				default:
					Layout.DrawText(option, menuX+20, menuY+int32(i*40)+10, 20, textColor)
				}
			}
			Layout.EndDrawing()
		}
	}
	return g, table
}

func (g *Game) InitializeDndScreen() {
	Layout.BeginDrawing()
	menuOptions := []string{"Continue", "New Game", "Tutorial", "Exit"}
	selectedOption := 0
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		for !Layout.WindowShouldClose() {
			Layout.BeginDrawing() // Properly start drawing at the beginning of each frame
			pressed := Layout.GetKeyPressed()
			if Layout.IsKeyPressed(Layout.KeyDown) || pressed == 264 {
				selectedOption = (selectedOption + 1) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyUp) || pressed == 265 {
				selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyEnter) || pressed == 257 {
				switch menuOptions[selectedOption] {
				case "Exit":
					Layout.EndDrawing()
					return
				case "Continue":
					if g.GameState.DdayArea {
						Layout.EndDrawing()
						g.UserContinue()
						return
					}
				case "New Game":
					Layout.EndDrawing()
					g.dDay()
					return
				case "Tutorial":
					// TODO: Tutorial has not been implemented yet.
					// It will be implemented later on
					break
				}
			}
			// Count visible options for proper box scaling
			visibleOptions := 0
			for _, option := range menuOptions {
				if option == "Continue" && !g.GameState.DdayArea {
					continue
				}
				visibleOptions++
			}
			// Position menu on the right side as specified
			menuX := int32(Layout.GetScreenWidth())/2 - 100 + (int32(len(menuOptions) * 40))
			menuY := int32(Layout.GetScreenHeight()) / 2
			// Draw the menu box scaled to visible options
			col := Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A}
			Layout.DrawRectangle(menuX, menuY, 200, int32(visibleOptions*40), Layout.ColorAlpha(col, 0.7))
			// Draw menu options with proper positioning
			visibleIndex := 0
			for i, option := range menuOptions {
				if option == "Continue" && !g.GameState.DdayArea {
					continue
				}
				textColor := Layout.Color{R: Layout.White.R, G: Layout.White.G, B: Layout.White.B, A: Layout.White.A}
				if i == selectedOption {
					textColor = Layout.Color{R: Layout.Red.R, G: Layout.Red.G, B: Layout.Red.B, A: Layout.Red.A}
					col = Layout.Color{R: Layout.Gray.R, G: Layout.Gray.G, B: Layout.Gray.B, A: Layout.Gray.A}
					Layout.DrawRectangle(menuX, menuY+int32(visibleIndex*40), 200, 40, Layout.ColorAlpha(col, 0.3))
				}
				Layout.DrawText(option, menuX+20, menuY+int32(visibleIndex*40)+10, 20, textColor)
				visibleIndex++
			}
			Layout.EndDrawing() // End drawing at the end of each frame
		}
	}
}

func (g *Game) InitializeCoopScreen() {
	Layout.BeginDrawing()
	menuOptions := []string{"Campaign", "Dnd", "PVP", "Exit"}
	selectedOption := 0
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		for !Layout.WindowShouldClose() {
			pressed := Layout.GetKeyPressed()
			if Layout.IsKeyPressed(Layout.KeyDown) || pressed == 264 {
				selectedOption = (selectedOption + 1) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyUp) || pressed == 265 {
				selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyEnter) || pressed == 257 {
				switch menuOptions[selectedOption] {
				case "Exit":
					Layout.ClearBackground(Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A})
					return
				case "Continue":
					g.Options.DndMode = 0
				case "New Game":
					g.Options.GameMode = 0
				}
			}
			// Get the dimensions for the Black Box
			menuX := int32(Layout.GetScreenWidth())/2 - 100 + (int32(len(menuOptions) * 40))
			menuY := int32(Layout.GetScreenHeight()) / 2
			// Draw the menu box
			col := Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A}
			Layout.DrawRectangle(menuX, menuY, 200, int32(len(menuOptions)*40), Layout.ColorAlpha(col, 0.7))
			// Draw menu options
			for i, option := range menuOptions {
				textColor := Layout.Color{R: Layout.White.R, G: Layout.White.G, B: Layout.White.B, A: Layout.White.A}
				if i == selectedOption {
					textColor = Layout.Color{R: Layout.Red.R, G: Layout.Red.G, B: Layout.Red.B, A: Layout.Red.A}
					col = Layout.Color{R: Layout.Gray.R, G: Layout.Gray.G, B: Layout.Gray.B, A: Layout.Gray.A}
					Layout.DrawRectangle(menuX, menuY+int32(i*40), 200, 40, Layout.ColorAlpha(col, 0.3))
				}
				if option == "Dnd" && g.Options.DndMode == 0 {
					continue
				} else {
					Layout.DrawText(option, menuX+20, menuY+int32(i*40)+10, 20, textColor)
				}
			}
			Layout.EndDrawing()
			Layout.BeginDrawing() // Begin new frame for next iteration
		}
		Layout.EndDrawing() // Final end drawing call
	}
}
