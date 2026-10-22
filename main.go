package main

import (
	"example.com/academy/catalog"
	"fmt"
)

type Valuable interface{
	Total() int
}


func reportTotal(items []Valuable) int {
	total := 0
	for i:=0 ;i < len(items);i++{
		total += items[i].Total()
	}
	return total
}



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

	items := []Valuable{
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
	fmt.Println(items[0].Total())
	fmt.Println(items[1].Total())
	fmt.Println(items[2].Total())
	fmt.Println(reportTotal(items))

}
