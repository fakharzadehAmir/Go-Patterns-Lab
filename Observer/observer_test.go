package Observer

import (
	"testing"
)

func TestObserverPattern(t *testing.T) {
	// Create a new product
	product := Product{ItemName: "Product A", subscribers: make(map[int64]*ISubscriber)}

	// Create a customer and a company observer
	NewCostumer(product)
	NewCompany(product)

	// Ensure that the subscribers are correctly subscribed to the product
	if len(product.subscribers) != 2 {
		t.Errorf("Expected 2 subscribers, got %d", len(product.subscribers))
	}

	// Notify observers about the availability of the product
	product.NotifySubs()
}
