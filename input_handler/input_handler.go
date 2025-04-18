package input_handler

// Module that contains functions that control the user's movement, reloading, and interacting with the vehicles

import (
	// Modules that are used for testing
	"fmt"

	term "github.com/nsf/termbox-go"
	// My own Modules
	// Use the dnd module instead and use the layout module as well.
	// We just need the dnd.Player instance inside a function parameter
	// We also need the functions that will operate the game which is in the layout module
	game "github.com/alphastigma101/Coconuts-At-Wars/game"
)

const (
	// Support for controllers
	BUTTON_X = "X"
	BUTTON_Y = "Y"
	BUTTON_A = "A"
	BUTTON_B = "B"
	// Support for Keyboards
	//KeyA = "a,A"
	//KeyW = "w,W"
	//KeyD = "d,D"
	KeyA = "a"
	KeyW = "w"
	KeyD = "d"
)

type Command interface {
	Execute(actor *game.Actor)
}

type InputHandler struct {
	ButtonX Command
	ButtonY Command
	ButtonA Command
	ButtonB Command
	KeyA    Command
	KeyW    Command
	KeyD    Command
}

type JumpCommand struct{}

func (j JumpCommand) jump(actor *game.Actor) {

}

func (j *JumpCommand) Execute(actor *game.Actor) {
	j.jump(actor)
}

// Struct that represents the fire command
type FireCommand struct{}

func (f FireCommand) fireGun(actor *game.Actor) {

}

func (f *FireCommand) Execute(actor *game.Actor) {
	f.fireGun(actor)
}

type DuckCommand struct{}

func (d DuckCommand) duck(actor *game.Actor) {

}

func (d *DuckCommand) Execute(actor *game.Actor) {
	d.duck(actor)
}

type ReloadCommand struct{}

func (r ReloadCommand) reload(actor *game.Actor) {

}

func (r *ReloadCommand) Execute(actor *game.Actor) {
	r.reload(actor)
}

func isPressed(hotkey string, input string) bool {
	//var split_input []string
	//if strings.Contains(input, ",") {
	//split_input = strings.Split(input, ",")
	//}
	if input == "w" && input == hotkey {
		fmt.Println("w pressed")
		return true
	} else if input == "a" && input == hotkey {
		fmt.Println("a pressed")
		return true
	} else if input == "d" && input == hotkey {
		fmt.Println("d pressed")
		return true
	} else if input == "X" && input == hotkey {
		fmt.Println("X pressed")
		return true
	} else if input == "A" && input == hotkey {
		fmt.Println("A pressed")
		return true
	} else if input == "B" && input == hotkey {
		fmt.Println("B pressed")
		return true
	} else if input == "Y" && input == hotkey {
		fmt.Println("Y pressed")
		return true
	}
	return false
}

// Function that handles user input
// Testing with termbox and will swap over to glxfw instead 2/14/2025
func HandleInput(h *InputHandler) Command {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	switch ev := term.PollEvent(); ev.Type {
	case term.EventKey:
		if isPressed(BUTTON_X, string(ev.Ch)) {
			return h.ButtonX
		} else if isPressed(BUTTON_Y, string(ev.Ch)) {
			return h.ButtonY
		} else if isPressed(BUTTON_A, string(ev.Ch)) {
			return h.ButtonA
		} else if isPressed(BUTTON_B, string(ev.Ch)) {
			return h.ButtonB
		} else if isPressed(KeyA, string(ev.Ch)) {
			return h.KeyA
		} else if isPressed(KeyW, string(ev.Ch)) {
			return h.KeyW
		} else if isPressed(KeyD, string(ev.Ch)) {
			return h.KeyD
		} else if isPressed(BUTTON_B, string(ev.Ch)) {
			return h.ButtonB
		}
	}
	return nil
}
