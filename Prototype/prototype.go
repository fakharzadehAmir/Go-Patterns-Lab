package Prototype

import "fmt"

// IShape is an interface representing a shape that can be cloned.
type IShape interface {
	Clone() IShape
	Print()
}

// Circle represents a circular shape with a radius.
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle instance with the given radius.
func NewCircle(radius float64) IShape {
	return &Circle{
		Radius: radius,
	}
}

// Clone creates a new Circle with the same radius.
func (c *Circle) Clone() IShape {
	return NewCircle(c.Radius)
}

// Print prints information about the Circle.
func (c *Circle) Print() {
	fmt.Printf("This is a Circle with radius %f", c.Radius)
}

// Rectangle represents a rectangular shape with height and width.
type Rectangle struct {
	Height float64
	Width  float64
}

// NewRectangle creates a new Rectangle instance with the given height and width.
func NewRectangle(height, width float64) IShape {
	return &Rectangle{
		Height: height,
		Width:  width,
	}
}

// Clone creates a new Rectangle with the same height and width.
func (r *Rectangle) Clone() IShape {
	return NewRectangle(r.Height, r.Width)
}

// Print prints information about the Rectangle.
func (r *Rectangle) Print() {
	fmt.Printf("This is a Rectangle with height %f and width %f", r.Height, r.Width)
}
