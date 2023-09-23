package State

import (
	"testing"
)

// TestCompleteTrafficLightCycle tests a complete cycle of traffic light state transitions.
func TestCompleteTrafficLightCycle(t *testing.T) {
	trafficLight := &TrafficLight{}

	// Initially, the traffic light is green.
	trafficLight.setState("Green")

	// Test the initial state (Green).
	if trafficLight.state.Transition() != "Green -> Yellow" {
		t.Errorf("Expected Green -> Yellow transition, got %s", trafficLight.state.Transition())
	}
	if trafficLight.state.MovementStatus() != "Now you can go!" {
		t.Errorf("Expected 'Now you can go!' status, got %s", trafficLight.state.MovementStatus())
	}

	// Transition to Yellow.
	trafficLight.setState("Yellow")

	// Test the Yellow state.
	if trafficLight.state.Transition() != "Yellow -> Green" {
		t.Errorf("Expected Yellow -> Green transition, got %s", trafficLight.state.Transition())
	}
	if trafficLight.state.MovementStatus() != "Be ready to go!" {
		t.Errorf("Expected 'Be ready to go!' status, got %s", trafficLight.state.MovementStatus())
	}

	// Transition to Red.
	trafficLight.setState("Red")

	// Test the Red state.
	if trafficLight.state.Transition() != "Red -> Green" {
		t.Errorf("Expected Red -> Green transition, got %s", trafficLight.state.Transition())
	}
	if trafficLight.state.MovementStatus() != "You have to stop, while the light is red" {
		t.Errorf("Expected 'You have to stop, while the light is red' status, got %s", trafficLight.state.MovementStatus())
	}
}
