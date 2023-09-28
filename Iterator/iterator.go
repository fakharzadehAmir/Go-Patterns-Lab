package Iterator

// Iterator interface defines methods hasNext() and next() to iterate through a collection.
type Iterator interface {
	hasNext() bool     // Checks if there are more items in the collection.
	next() interface{} // Returns the next item in the collection.
}

// MenuIterator is a concrete iterator that implements the Iterator interface.
type MenuIterator struct {
	position  int
	menuItems []MenuItem
}

// NewMenuIterator creates a new MenuIterator instance with a collection of menu items.
func NewMenuIterator(menuItems []MenuItem) *MenuIterator {
	return &MenuIterator{menuItems: menuItems}
}

// hasNext checks if there are more items to iterate through.
func (mi *MenuIterator) hasNext() bool {
	if len(mi.menuItems) > mi.position {
		return true
	}
	return false
}

// next returns the next menu item and advances the iterator position.
func (mi *MenuIterator) next() interface{} {
	if mi.hasNext() {
		item := mi.menuItems[mi.position]
		mi.position++
		return item
	}
	return nil
}

// IAggregate interface defines a method createIterator() to create an Iterator instance.
type IAggregate interface {
	createIterator() Iterator
}

// Menu is a concrete aggregate that holds a collection of menu items.
type Menu struct {
	items []MenuItem
}

// NewMenu creates a new Menu instance with an initial collection of menu items.
func NewMenu(newItems []MenuItem) *Menu {
	return &Menu{items: newItems}
}

// AddItems appends a new menu item to the collection.
func (m *Menu) AddItems(newItem MenuItem) {
	m.items = append(m.items, newItem)
}

// createIterator creates a MenuIterator for iterating through the menu items.
func (m *Menu) createIterator() Iterator {
	return NewMenuIterator(m.items)
}

// MenuItem represents a menu item with attributes such as name, description, and price.
type MenuItem struct {
	Name        string
	Description string
	Price       int64
}
