package catalog

type Product struct {
	Name string
	Price int
	Quantity int
}

func (p Product) Total() int{
	return p.Price * p.Quantity
}