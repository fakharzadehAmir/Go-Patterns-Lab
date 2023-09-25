package Observer

import "fmt"

// ISubscriber defines the interface that subscribers (observers) must implement.
type ISubscriber interface {
	notify()
	getID() int64
}

// Product represents the subject that observers subscribe to.
type Product struct {
	ItemName    string
	subscribers map[int64]*ISubscriber
}

// Subscribe allows subscribers to subscribe to the product.
func (p *Product) Subscribe(subscriber ISubscriber) {
	p.subscribers[subscriber.getID()] = &subscriber
}

// Unsubscribe allows subscribers to unsubscribe from the product.
func (p *Product) Unsubscribe(subscriber ISubscriber) {
	delete(p.subscribers, subscriber.getID())
}

// NotifySubs notifies all subscribed observers when a change occurs in the product.
func (p *Product) NotifySubs() {
	for _, sub := range p.subscribers {
		(*sub).notify()
	}
}

// Costumer represents a customer observer.
type Costumer struct {
	id         int64
	pubProduct Product
}

// NewCostumer creates a new customer observer and subscribes it to a product.
func NewCostumer(product Product) *Costumer {
	newID := int64(len(product.subscribers))
	newCostumer := &Costumer{id: newID, pubProduct: product}
	product.Subscribe(newCostumer)
	return newCostumer
}

// notify notifies the customer when the product becomes available.
func (c *Costumer) notify() {
	fmt.Printf("Customer (%v), %v is available now\n", c.id, c.pubProduct.ItemName)
}

// getID returns the ID of the customer.
func (c *Costumer) getID() int64 {
	return c.id
}

// Company represents a company observer.
type Company struct {
	id         int64
	pubProduct Product
}

// NewCompany creates a new company observer and subscribes it to a product.
func NewCompany(product Product) *Company {
	newID := int64(len(product.subscribers))
	newCompany := &Company{id: newID, pubProduct: product}
	product.Subscribe(newCompany)
	return newCompany
}

// notify notifies the company when the product becomes available.
func (c *Company) notify() {
	fmt.Printf("Company (%v), %v is available now\n", c.id, c.pubProduct.ItemName)
}

// getID returns the ID of the company.
func (c *Company) getID() int64 {
	return c.id
}
