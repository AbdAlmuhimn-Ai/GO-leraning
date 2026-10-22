package catalog

import "errors"

type Product struct {
	Name     string
	Price    int
	Quantity int
}

func (p Product) Total() int {
	return p.Price * p.Quantity
}

func (p *Product) Sell(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	if amount > p.Quantity {
		return errors.New("insufficient stock")
	}

	p.Quantity -= amount
	return nil
}

func (p *Product) Restock(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	p.Quantity += amount
	return nil
}
