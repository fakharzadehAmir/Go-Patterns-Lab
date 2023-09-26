package Command

import "fmt"

// ILightCommand is the interface that defines the execute method
// for all concrete command objects.
type ILightCommand interface {
	execute()
}

// Light represents the receiver in the Command pattern. It's the
// object that the commands will act upon.
type Light struct {
	isOn bool
}

// On turns the light on and prints a message.
func (l *Light) On() {
	l.isOn = true
	fmt.Printf("Light has been turned on!")
}

// Off turns the light off and prints a message.
func (l *Light) Off() {
	l.isOn = false
	fmt.Printf("Light has been turned off!")
}

// ChangeBright changes the brightness of the light and prints a message.
func (l *Light) ChangeBright() {
	fmt.Printf("Brightness has been changed!")
}

// TurnOnCommand is a concrete command that implements the ILightCommand
// interface to turn the light on.
type TurnOnCommand struct {
	light *Light
}

// execute method of TurnOnCommand turns the light on by calling the
// On method of the Light object.
func (on *TurnOnCommand) execute() {
	on.light.On()
}

// TurnOffCommand is a concrete command that implements the ILightCommand
// interface to turn the light off.
type TurnOffCommand struct {
	light *Light
}

// execute method of TurnOffCommand turns the light off by calling the
// Off method of the Light object.
func (off *TurnOffCommand) execute() {
	off.light.Off()
}

// ChangeBrightCommand is a concrete command that implements the ILightCommand
// interface to change the brightness of the light.
type ChangeBrightCommand struct {
	light *Light
}

// execute method of ChangeBrightCommand changes the brightness of the light
// by calling the ChangeBright method of the Light object.
func (cb *ChangeBrightCommand) execute() {
	cb.light.ChangeBright()
}
