/*
This is the game module
*/
package game

// This Module holds the Game struct which consists of nested structs that keeps track of the user's data
import (
	// My modules
	dnd "github.com/alphastigma101/Coconuts-At-Wars/Dnd"
	coop "github.com/alphastigma101/Coconuts-At-Wars/cooperative"
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
	"github.com/alphastigma101/Coconuts-At-Wars/main_game"
	"github.com/alphastigma101/Coconuts-At-Wars/options"
)

type DataBaseProperties Layout.Properties

// During the campaign, if the user clicks a certain button, it will pop up a menu
// They can either save, exit, load, and connect
type GameOptions interface {
	Options() interface{}
}

// Struct that will keep track of the damage and the position of the player
type gameActor struct {
	Health   int
	Position Layout.Position
}

type Actor gameActor

// Layout of the game. It will copied to the database assigned with a unique player id
// It also will be the options the user will see once they get pass the titlescreen
type Game struct {
	Options     *options.Options // Options that the user can configure
	Campaign    *main_game.MainGame
	Dnd         *dnd.Dnd
	Cooperative *coop.Cooperative
	Players     *dnd.Player
	GameActor   *Actor
}

func (g *Game) InitializeCampaignScreen() {
	// Reuse the main menu box to draw out the campaign screen:
	//g.Options.Game2D.(Layout.Game2D).GameLayout.MainMenu.Rectangle
	// Should have Continue, StartOver, Exit
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		// Everything needs to be unloaded such as images, textures, etc
		//g.Options.Game2D.InitializeMainMenuScreen(g) // If user clicks exit, go back to main menu screen
	}
	panic("Campaign is being executed!.Not Implemented Yet!")
}

func (g *Game) InitializeOptionsScreen(table *Layout.Table) (interface{}, interface{}) {
	Layout.BeginDrawing()
	menuOptions := []string{"Dnd Option", "Game Mode", "Exit"}
	selectedOption := 0
	if g.Options.GameMode == 1 {
		// Do stuff for 3D
	} else {
		exitMenu := false
		for !exitMenu && !Layout.WindowShouldClose() {
			pressed := Layout.GetKeyPressed()
			// Handle events and input
			if Layout.IsKeyPressed(Layout.KeyDown) || pressed == 264 {
				selectedOption = (selectedOption + 1) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyUp) || pressed == 265 {
				selectedOption = (selectedOption - 1 + len(menuOptions)) % len(menuOptions)
			} else if Layout.IsKeyPressed(Layout.KeyEnter) || pressed == 257 {
				// Process selection
				switch menuOptions[selectedOption] {
				case "Exit":
					exitMenu = true
					Layout.ClearBackground(Layout.Color{R: Layout.Black.R, G: Layout.Black.G, B: Layout.Black.B, A: Layout.Black.A})
					g.Options.Game2D.InitializeMainMenuScreen(g, table) // If user clicks exit, go back to main menu screen
				}
			} else if Layout.IsKeyPressed(Layout.KeyLeft) || pressed == 263 {
				// Handle left arrow for decreasing value
				switch menuOptions[selectedOption] {
				case "Dnd Option":
					g.Options.DndMode = 0 // Turn off DnD
				case "Game Mode":
					g.Options.GameMode = 0 // Set to 2D mode
				}
			} else if Layout.IsKeyPressed(Layout.KeyRight) || pressed == 262 {
				// Handle right arrow for increasing value
				switch menuOptions[selectedOption] {
				case "Dnd Option":
					g.Options.DndMode = 1 // Turn on DnD
					//table.Options.Update(g)
				case "Game Mode":
					g.Options.GameMode = 1 // Set to 3D mode
					//table.Options.Update(g)
				}
			}
			// Draw menu box
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

				// Draw option text with toggle indicators for DnD and Game Mode
				switch option {
				case "Dnd Option":
					dndStatus := "Off"
					if g.Options.DndMode == 1 {
						dndStatus = "On"
					}
					Layout.DrawText(option+": < "+dndStatus+" >", menuX+20, menuY+int32(i*40)+10, 20, textColor)
				case "Game Mode":
					modeStatus := "2D"
					if g.Options.GameMode == 1 {
						modeStatus = "3D"
					}
					Layout.DrawText(option+": < "+modeStatus+" >", menuX+20, menuY+int32(i*40)+10, 20, textColor)
				default:
					Layout.DrawText(option, menuX+20, menuY+int32(i*40)+10, 20, textColor)
				}
			}
			Layout.EndDrawing()
			Layout.BeginDrawing() // Begin new frame for next iteration
		}
		Layout.EndDrawing() // Final end drawing call
	}
	return g, table
}
