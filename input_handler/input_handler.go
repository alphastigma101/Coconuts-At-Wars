type Command interface {
	execute()
}

type JumpCommand struct {
    // fields go here if needed
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
