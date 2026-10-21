package main

import (
	"example.com/academy/catalog"
	"fmt"
)

func main() {
	p :=catalog.Product{
		Name:     "Mouse",
		Price:    20,
		Quantity: 3,
	}
	fmt.Println(p.Total())
	err := p.Sell(2)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Sold")
		fmt.Println(p.Quantity)
	}
	err = p.Sell(5)
	if err != nil {
		fmt.Println(err)
		fmt.Println(p.Quantity)
	}

	err = p.Restock(0)
	if err != nil {
		fmt.Println(err)
	}

	err = p.Restock(4)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Restocked")
		fmt.Println(p.Quantity)
		fmt.Println(p.Total())
	}
}
