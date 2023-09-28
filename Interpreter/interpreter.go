package Interpreter

// Expression is the interface that all expressions must implement.
type Expression interface {
	interpret() int // interpret method should evaluate the expression and return an integer result.
}

// Number represents a terminal expression that holds a numerical value.
type Number struct {
	value int
}

func (n *Number) interpret() int {
	return n.value
}

// Operation represents a non-terminal expression that combines two sub-expressions with an operator.
type Operation struct {
	leftExp  Expression
	rightExp Expression
	operator string // The operator, such as "+" or "-".
}

func (o *Operation) interpret() int {
	if o.operator == "+" {
		return o.leftExp.interpret() + o.rightExp.interpret()
	} else if o.operator == "-" {
		return o.leftExp.interpret() - o.rightExp.interpret()
	}
	return 0
}
