package catalog

import "errors"

func FindProductIndex(products []Product, name string) int {
	for i := range products {
		if products[i].Name == name {
			return i
		}
	}
	return -1
}
func SellByName(products []Product, name string, amount int) error {
	for i := range products {
		if products[i].Name == name {
			return products[i].Sell(amount)
		}
	}
	return errors.New("product not found")
}