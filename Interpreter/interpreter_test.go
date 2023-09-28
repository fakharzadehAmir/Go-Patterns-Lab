package Interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	// Create terminal expressions
	num1 := &Number{value: 3}
	num2 := &Number{value: 5}
	num3 := &Number{value: 2}

	// Create non-terminal expressions for individual operations
	sum1 := &Operation{
		leftExp:  num1,
		rightExp: num2,
		operator: "+",
	}

	subtract := &Operation{
		leftExp:  sum1, // Use the result of the previous sum (3 + 5)
		rightExp: num3,
		operator: "-",
	}

	// Test interpreting the combined expression (3 + 5 - 2)
	result := subtract.interpret()
	expected := 6
	if result != expected {
		t.Errorf("Expected result for combined expression: %d, got: %d", expected, result)
	}
}
