package Visitor

import "fmt"

// ICompany is the element interface that declares the accept method for visitors.
type ICompany interface {
	accept(visitor Visitor)
}

// NationalBank is a concrete element representing a national bank.
type NationalBank struct {
	foundedYear int
}

// accept allows the visitor to visit the national bank and perform an operation.
func (nb *NationalBank) accept(visitor Visitor) {
	visitor.visitBank(nb)
}

// ChocolateFactory is a concrete element representing a chocolate factory.
type ChocolateFactory struct {
	foundedYear int
}

// accept allows the visitor to visit the chocolate factory and perform an operation.
func (cf *ChocolateFactory) accept(visitor Visitor) {
	visitor.visitChocolateFactory(cf)
}

// AppleInc is a concrete element representing Apple Inc.
type AppleInc struct {
	foundedYear int
}

// accept allows the visitor to visit Apple Inc. and perform an operation.
func (ai *AppleInc) accept(visitor Visitor) {
	visitor.visitApple(ai)
}

// Visitor is the visitor interface that declares visit methods for concrete elements.
type Visitor interface {
	visitBank(bank *NationalBank)
	visitChocolateFactory(chocoFac *ChocolateFactory)
	visitApple(apple *AppleInc)
}

// InsuranceVisitor is a concrete visitor that calculates insurance costs for different companies.
type InsuranceVisitor struct{}

// visitBank calculates insurance cost for a national bank.
func (iv *InsuranceVisitor) visitBank(bank *NationalBank) {
	fmt.Printf("Insurance cost for a bank founded on %v: $10,000", bank.foundedYear)
}

// visitChocolateFactory calculates insurance cost for a chocolate factory.
func (iv *InsuranceVisitor) visitChocolateFactory(chocoFac *ChocolateFactory) {
	fmt.Printf("Insurance cost for a chocolate factory founded on %v: $10,000", chocoFac.foundedYear)
}

// visitApple calculates insurance cost for Apple Inc.
func (iv *InsuranceVisitor) visitApple(apple *AppleInc) {
	fmt.Printf("Insurance cost for Apple Inc. founded on %v: $5,000", apple.foundedYear)
}
