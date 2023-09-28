package Iterator

import (
	"testing"
)

func TestMenuIterator(t *testing.T) {
	// Create some menu items
	item1 := MenuItem{Name: "Burger", Description: "Delicious burger", Price: 100}
	item2 := MenuItem{Name: "Pizza", Description: "Tasty pizza", Price: 150}
	item3 := MenuItem{Name: "Salad", Description: "Healthy salad", Price: 70}

	// Create a menu and add items to it
	menu := NewMenu([]MenuItem{item1, item2, item3})

	// Create an iterator for the menu
	iterator := menu.createIterator()

	// Iterate through the menu items and check if they match the expected values
	expectedItems := []MenuItem{item1, item2, item3}
	index := 0

	for iterator.hasNext() {
		item := iterator.next().(MenuItem)
		if item != expectedItems[index] {
			t.Errorf("Expected item: %v, got: %v", expectedItems[index], item)
		}
		index++
	}

	// Check if the iterator reaches the end of the collection
	if iterator.hasNext() {
		t.Error("Iterator should reach the end of the collection.")
	}
}

func TestMenuAddItems(t *testing.T) {
	// Create an empty menu
	menu := NewMenu([]MenuItem{})

	// Add a new menu item
	newItem := MenuItem{Name: "Soda", Description: "Refreshing soda", Price: 20}
	menu.AddItems(newItem)

	// Create an iterator for the menu
	iterator := menu.createIterator()

	// Check if the added item is present in the menu
	if !iterator.hasNext() {
		t.Error("Iterator should have at least one item after adding.")
	}

	item := iterator.next().(MenuItem)
	if item != newItem {
		t.Errorf("Expected item: %v, got: %v", newItem, item)
	}
}
