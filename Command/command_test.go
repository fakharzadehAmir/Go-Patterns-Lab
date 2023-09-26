package Command

import (
	"testing"
)

func TestLightCommands(t *testing.T) {
	// Create a new Light instance
	light := &Light{}

	// Create command instances
	turnOnCommand := &TurnOnCommand{light}
	turnOffCommand := &TurnOffCommand{light}
	changeBrightCommand := &ChangeBrightCommand{light}

	// Execute the commands and test the output
	t.Run("Test TurnOnCommand", func(t *testing.T) {
		turnOnCommand.execute()
		if !light.isOn {
			t.Errorf("Expected light to be on after TurnOnCommand")
		}
	})

	t.Run("Test TurnOffCommand", func(t *testing.T) {
		turnOffCommand.execute()
		if light.isOn {
			t.Errorf("Expected light to be off after TurnOffCommand")
		}
	})

	t.Run("Test ChangeBrightCommand", func(t *testing.T) {
		changeBrightCommand.execute()
		// Since ChangeBrightCommand doesn't change the state of 'isOn',
		// we don't need to check it here.
	})

	// Additional test for the Light methods
	t.Run("Test Light Methods", func(t *testing.T) {
		// Test turning the light on
		light.On()
		if !light.isOn {
			t.Errorf("Expected light to be on after calling On()")
		}

		// Test turning the light off
		light.Off()
		if light.isOn {
			t.Errorf("Expected light to be off after calling Off()")
		}

		// Test changing brightness
		light.ChangeBright()
		// No state change is expected for brightness change, so we don't need to check it here.
	})
}
