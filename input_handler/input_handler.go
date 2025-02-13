package input_handler
const (
    BUTTON_X = "X"
    BUTTON_Y = "Y"
    BUTTON_A = "A"
    BUTTON_B = "B"
)

type Command interface {
	execute()
}

type InputHandler struct {
    buttonX Command
    buttonY Command
    buttonA Command
    buttonB Command
}

type JumpCommand struct {
  
}

func (j JumpCommand) jump() {

}

func (j JumpCommand) execute() {
    j.jump()
}

type FireCommand struct {
    
}

func (f FireCommand) fireGun() {

}

func (f FireCommand) execute() {
    f.fireGun()
}

func (h *InputHandler) HandleInput() {
    if isPressed(BUTTON_X) {
        h.buttonX.execute()
    } else if isPressed(BUTTON_Y) {
        h.buttonY.execute()
    } else if isPressed(BUTTON_A) {
        h.buttonA.execute()
    } else if isPressed(BUTTON_B) {
        h.buttonB.execute()
    }
}