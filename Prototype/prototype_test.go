package Prototype

import (
	"testing"
)

func TestShapeCloning(t *testing.T) {
	// Create a circle and a rectangle
	circle := NewCircle(5.0)
	rectangle := NewRectangle(4.0, 6.0)

	// Clone the circle and rectangle
	cloneCircle := circle.Clone()
	cloneRectangle := rectangle.Clone()

	// Check if the cloned shapes are of the correct type
	if _, ok := cloneCircle.(*Circle); !ok {
		t.Errorf("Expected cloneCircle to be of type *Circle, got %T", cloneCircle)
	}

	if _, ok := cloneRectangle.(*Rectangle); !ok {
		t.Errorf("Expected cloneRectangle to be of type *Rectangle, got %T", cloneRectangle)
	}

	// Check if the cloned shapes have the same properties as the originals
	if cloneCircle.(*Circle).Radius != circle.(*Circle).Radius {
		t.Errorf("Expected cloneCircle to have radius %f, got %f", circle.(*Circle).Radius, cloneCircle.(*Circle).Radius)
	}

	if cloneRectangle.(*Rectangle).Height != rectangle.(*Rectangle).Height || cloneRectangle.(*Rectangle).Width != rectangle.(*Rectangle).Width {
		t.Errorf("Expected cloneRectangle to have height %f and width %f, got height %f and width %f",
			rectangle.(*Rectangle).Height, rectangle.(*Rectangle).Width,
			cloneRectangle.(*Rectangle).Height, cloneRectangle.(*Rectangle).Width)
	}
}

func TestShapePrinting(t *testing.T) {
	// Create a circle and a rectangle
	circle := NewCircle(3.0)
	rectangle := NewRectangle(2.0, 4.0)

	// Create a clone of the circle and a clone of the rectangle
	cloneCircle := circle.Clone()
	cloneRectangle := rectangle.Clone()

	// Create a slice of shapes
	shapes := []IShape{circle, cloneCircle, rectangle, cloneRectangle}

	// Print information about each shape
	for i, shape := range shapes {
		t.Logf("Shape %d:", i+1)
		shape.Print()
		t.Logf("\n")
	}
}
