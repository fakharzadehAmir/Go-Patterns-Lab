package State

import "fmt"

// IState represents the interface for different traffic light states.
type IState interface {
	Transition() string     // Method for transitioning to the next state
	MovementStatus() string // Method for describing the movement status in the current state
}

// GreenLight represents the Green traffic light state.
type GreenLight struct{}

func (gl GreenLight) Transition() string {
	return fmt.Sprintf("Green -> Yellow")
}

func (gl GreenLight) MovementStatus() string {
	return fmt.Sprintf("Now you can go!")
}

// YellowLight represents the Yellow traffic light state.
type YellowLight struct{}

func (yl YellowLight) Transition() string {
	return fmt.Sprintf("Yellow -> Green")
}

func (yl YellowLight) MovementStatus() string {
	return fmt.Sprintf("Be ready to go!")
}

// RedLight represents the Red traffic light state.
type RedLight struct{}

func (rl RedLight) Transition() string {
	return fmt.Sprintf("Red -> Green")
}

func (rl RedLight) MovementStatus() string {
	return fmt.Sprintf("You have to stop, while the light is red")
}

// TrafficLight represents the context that holds the current state.
type TrafficLight struct {
	state IState
}

// setState sets the state of the traffic light based on the given status.
func (tl *TrafficLight) setState(status string) {
	if status == "Red" {
		tl.state = &RedLight{}
	} else if status == "Yellow" {
		tl.state = &YellowLight{}
	} else if status == "Green" {
		tl.state = &GreenLight{}
	}
	return
}
