package Builder

// Burger represents a customizable burger.
type Burger struct {
	Patty    string
	Toppings []string
	Cheese   string
}

// IBurgerBuilder is the interface for building custom burgers.
type IBurgerBuilder interface {
	AddToppings(toppings []string)
	AddCheese(cheese string)
	SetPatty()
	Cook() *Burger
}

// getBuilder returns a specific burger builder based on the type.
func getBuilder(typeBuilder string) IBurgerBuilder {
	if typeBuilder == "hamburger" {
		return &HamburgerBuilder{}
	} else if typeBuilder == "chicken" {
		return &ChickenBuilder{}
	}
	return nil
}

// HamburgerBuilder is a concrete builder for creating hamburger burgers.
type HamburgerBuilder struct {
	toppings []string
	patty    string
	cheese   string
}

// AddCheese adds cheese to the hamburger.
func (hb *HamburgerBuilder) AddCheese(cheese string) {
	hb.cheese = cheese
}

// AddToppings adds toppings to the hamburger.
func (hb *HamburgerBuilder) AddToppings(toppings []string) {
	hb.toppings = toppings
}

// SetPatty sets the patty type to beef for the hamburger.
func (hb *HamburgerBuilder) SetPatty() {
	hb.patty = "beef"
}

// Cook builds and returns the custom hamburger.
func (hb *HamburgerBuilder) Cook() *Burger {
	return &Burger{
		Patty:    hb.patty,
		Cheese:   hb.cheese,
		Toppings: hb.toppings,
	}
}

// ChickenBuilder is a concrete builder for creating chicken burgers.
type ChickenBuilder struct {
	toppings []string
	patty    string
	cheese   string
}

// AddCheese adds cheese to the chicken burger.
func (cb *ChickenBuilder) AddCheese(cheese string) {
	cb.cheese = cheese
}

// AddToppings adds toppings to the chicken burger.
func (cb *ChickenBuilder) AddToppings(toppings []string) {
	cb.toppings = toppings
}

// SetPatty sets the patty type to chicken for the chicken burger.
func (cb *ChickenBuilder) SetPatty() {
	cb.patty = "chicken"
}

// Cook builds and returns the custom chicken burger.
func (cb *ChickenBuilder) Cook() *Burger {
	return &Burger{
		Patty:    cb.patty,
		Cheese:   cb.cheese,
		Toppings: cb.toppings,
	}
}
