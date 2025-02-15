package input_handler

import (
	// Modules that are used for testing
	"fmt"

	term "github.com/nsf/termbox-go"
)

type GameActor struct {
	Health int
}

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
	Execute(actor *GameActor)
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

func (j JumpCommand) jump(actor *GameActor) {

}

func (j *JumpCommand) Execute(actor *GameActor) {
	j.jump(actor)
}

// Struct that represents the fire command
type FireCommand struct{}

func (f FireCommand) fireGun(actor *GameActor) {

}

func (f *FireCommand) Execute(actor *GameActor) {
	f.fireGun(actor)
}

type DuckCommand struct{}

func (d DuckCommand) duck(actor *GameActor) {

}

func (d *DuckCommand) Execute(actor *GameActor) {
	d.duck(actor)
}

type ReloadCommand struct{}

func (r ReloadCommand) reload(actor *GameActor) {

}

func (r *ReloadCommand) Execute(actor *GameActor) {
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
