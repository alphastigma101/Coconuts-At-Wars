package input_handler

import (
	// Modules that are used for testing
	"fmt"
	"strings"

	term "github.com/nsf/termbox-go"
)

type GameActor struct{}

const (
	// Support for controllers
	BUTTON_X = "X"
	BUTTON_Y = "Y"
	BUTTON_A = "A"
	BUTTON_B = "B"
	// Support for Keyboards
	KeyA = "a,A"
	KeyW = "w,W"
	KeyD = "d,D"
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

func isPressed(input string) bool {
	var split_input []string
	if strings.Contains(input, ",") {
		split_input = strings.Split(input, ",")
	}
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			if input == string(ev.Ch) || len(split_input) > 0 {
				if input == "w" || ((split_input[0] == "w") || split_input[1] == "W") {
					fmt.Println("w pressed")
					return true
				}
				if input == "a" || ((split_input[0] == "a") || split_input[1] == "A") {
					fmt.Println("a pressed")
					return true
				}
				if input == "d" || ((split_input[0] == "d") || split_input[1] == "D") {
					fmt.Println("d pressed")
					return true
				}
				if input == "q" || input == "Q" {
					// Quit on Q press
					return true
				}
				panic(ev.Err)
			}
		}
	}
}
func HandleInput(h *InputHandler) Command {
	if isPressed(BUTTON_X) {
		return h.ButtonX
	} else if isPressed(BUTTON_Y) {
		return h.ButtonY
	} else if isPressed(BUTTON_A) {
		return h.ButtonA
	} else if isPressed(BUTTON_B) {
		return h.ButtonB
	} else if isPressed(KeyA) {
		return h.KeyA
	} else if isPressed(KeyW) {
		return h.KeyW
	} else if isPressed(KeyD) {
		return h.KeyD
	} else if isPressed(BUTTON_B) {
		return h.ButtonB
	}
	return nil
}
