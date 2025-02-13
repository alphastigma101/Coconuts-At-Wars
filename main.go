package main
import (
    "fmt"
    "github.com/alphastigma101/Coconuts-At-Wars/input_handler"
)

func main() {
    // Get a greeting message and print it.
    message := "Hello"
    fmt.Println(message)
    handler := &input_handler.InputHandler{
        buttonX: &JumpCommand{},
        buttonY: &FireCommand{},
        buttonA: &DuckCommand{},
        buttonB: &ReloadCommand{},
    }
    
    // Call the handler in your game loop
    handler.HandleInput()
}