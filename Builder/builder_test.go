package Builder

import (
	"reflect"
	"testing"
)

func TestHamburgerBuilder(t *testing.T) {
	builder := getBuilder("hamburger")
	if builder == nil {
		t.Fatal("Expected HamburgerBuilder but got nil")
	}

	hamburgerBuilder, ok := builder.(*HamburgerBuilder)
	if !ok {
		t.Fatal("Expected HamburgerBuilder but got a different type")
	}

	// Build a hamburger with cheese and toppings
	hamburgerBuilder.SetPatty()
	hamburgerBuilder.AddCheese("cheddar")
	hamburgerBuilder.AddToppings([]string{"lettuce", "tomato"})

	burger := hamburgerBuilder.Cook()

	// Check if the burger is correctly built
	expectedBurger := &Burger{
		Patty:    "beef",
		Cheese:   "cheddar",
		Toppings: []string{"lettuce", "tomato"},
	}

	if !reflect.DeepEqual(burger, expectedBurger) {
		t.Errorf("Expected burger: %+v, but got: %+v", expectedBurger, burger)
	}
}

func TestChickenBuilder(t *testing.T) {
	builder := getBuilder("chicken")
	if builder == nil {
		t.Fatal("Expected ChickenBuilder but got nil")
	}

	chickenBuilder, ok := builder.(*ChickenBuilder)
	if !ok {
		t.Fatal("Expected ChickenBuilder but got a different type")
	}

	// Build a chicken burger with cheese and toppings
	chickenBuilder.SetPatty()
	chickenBuilder.AddCheese("swiss")
	chickenBuilder.AddToppings([]string{"lettuce", "onion"})

	burger := chickenBuilder.Cook()

	// Check if the burger is correctly built
	expectedBurger := &Burger{
		Patty:    "chicken",
		Cheese:   "swiss",
		Toppings: []string{"lettuce", "onion"},
	}

	if !reflect.DeepEqual(burger, expectedBurger) {
		t.Errorf("Expected burger: %+v, but got: %+v", expectedBurger, burger)
	}
}
