package main

import (
	"example.com/academy/catalog"
	"example.com/academy/report"
	"fmt"
)

func main() {

	p := catalog.Product{
		Name:     "Mouse",
		Price:    20,
		Quantity: 3,
	}
	fmt.Println(p.Total())
	err := p.Sell(2)
	if err != nil {
		fmt.Println(err)
	} else {
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
	} else {
		fmt.Println("Restocked")
		fmt.Println(p.Quantity)
		fmt.Println(p.Total())
	}

	items := []report.Valuable{
		p,
		catalog.Service{
			Name:     "Design",
			Rate:     40,
			Sessions: 2,
		},
		catalog.Product{
			Name:     "Keyboard",
			Price:    15,
			Quantity: 2,
		},
	}
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i].Total())
	}

	fmt.Println(report.Sum(items))
	fmt.Println(report.CountAbove(items, 50))

}
