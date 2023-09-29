package Memento

import "testing"

func TestMementoPattern(t *testing.T) {
	// Create an originator and caretaker
	originator := Originator{}
	caretaker := Caretaker{originator: &originator}

	// Write some text and create a snapshot (memento)
	caretaker.write("State 1")
	caretaker.write("State 2")
	caretaker.write("State 3")

	// Check if undo restores the previous state correctly
	caretaker.undo()
	if originator.text != "State 2" {
		t.Errorf("Expected text 'State 2' after undo, but got: %s", originator.text)
	}

	caretaker.undo()
	if originator.text != "State 1" {
		t.Errorf("Expected text 'State 1' after undo, but got: %s", originator.text)
	}

	// Create another snapshot
	caretaker.write("State 4")

	// Undo again to test redo functionality
	caretaker.undo()
	caretaker.undo()
	caretaker.undo()
	caretaker.undo()
	caretaker.undo()
	if originator.text != "" {
		t.Errorf("Expected text '' after redo, but got: %s", originator.text)
	}
}
