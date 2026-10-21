package main

import (
	"example.com/academy/catalog"
	"fmt"
)

func main() {
	p :=catalog.Product{
		Name:     "Mouse",
		Price:    20,
		Quantity: 2,
	}

	fmt.Println(p.Name, p.Total())
}
