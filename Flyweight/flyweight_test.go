package Flyweight

import (
	"testing"
)

func TestJerseyFactory_GetJerseyBySide(t *testing.T) {
	factory := NewJerseyFactory()

	homeJersey := factory.getJerseyBySide(homeTeam, "red")
	awayJersey := factory.getJerseyBySide(awayTeam, "blue")

	// Check if the jerseys are created and have the correct colors.
	if homeJersey == nil || homeJersey.getColor() != "red" {
		t.Error("Expected home jersey to be created with color 'red', but got nil or incorrect color.")
	}

	if awayJersey == nil || awayJersey.getColor() != "blue" {
		t.Error("Expected away jersey to be created with color 'blue', but got nil or incorrect color.")
	}

	// Repeated calls to getJerseyBySide should return the same instances.
	homeJersey2 := factory.getJerseyBySide(homeTeam, "green")
	if homeJersey != homeJersey2 {
		t.Error("Expected the same home jersey instance for repeated calls, but got different instances.")
	}
}

func TestPlayer(t *testing.T) {
	factory := NewJerseyFactory()

	// Create two players with different jersey colors.
	player1 := Player{
		Firstname: "John",
		Lastname:  "Doe",
		Overall:   85,
		Jersey:    factory.getJerseyBySide(homeTeam, "red"),
	}

	player2 := Player{
		Firstname: "Jane",
		Lastname:  "Smith",
		Overall:   90,
		Jersey:    factory.getJerseyBySide(awayTeam, "blue"),
	}

	// Check if the players have different jersey instances.
	if player1.Jersey == player2.Jersey {
		t.Error("Expected players to have different jersey instances, but they share the same instance.")
	}
}
