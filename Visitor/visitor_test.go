package Visitor

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestVisitorPattern(t *testing.T) {
	// Create concrete elements
	bank := &NationalBank{foundedYear: 1900}
	chocolateFactory := &ChocolateFactory{foundedYear: 2005}
	appleInc := &AppleInc{foundedYear: 1976}

	// Create a visitor
	insuranceVisitor := &InsuranceVisitor{}

	// Test visiting and calculating insurance costs
	bank.accept(insuranceVisitor)
	chocolateFactory.accept(insuranceVisitor)
	appleInc.accept(insuranceVisitor)

	// Ensure the correct insurance costs are calculated
	expectedBankOutput := "Insurance cost for a bank founded on 1900: $10,000"
	expectedChocoFacOutput := "Insurance cost for a chocolate factory founded on 2005: $10,000"
	expectedAppleOutput := "Insurance cost for Apple Inc. founded on 1976: $5,000"

	if got := captureOutput(func() { bank.accept(insuranceVisitor) }); got != expectedBankOutput {
		t.Errorf("Expected output for bank is %q, but got %q", expectedBankOutput, got)
	}

	if got := captureOutput(func() { chocolateFactory.accept(insuranceVisitor) }); got != expectedChocoFacOutput {
		t.Errorf("Expected output for chocolate factory is %q, but got %q", expectedChocoFacOutput, got)
	}

	if got := captureOutput(func() { appleInc.accept(insuranceVisitor) }); got != expectedAppleOutput {
		t.Errorf("Expected output for Apple Inc. is %q, but got %q", expectedAppleOutput, got)
	}
}

// Helper function to capture printed output
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	return string(out)
}
