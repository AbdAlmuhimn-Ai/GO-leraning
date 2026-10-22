package report

type Valuable interface {
	Value() int
}

func Sum(items []Valuable) int {
	sum := 0
	for i := 0; i < len(items); i++ {
		sum += items[i].Value()
	}
	return sum
}

func CountAbove(items []Valuable, threshold int) int {
	count := 0
	for i := 0; i < len(items); i++ {
		if items[i].Value() > threshold {
			count++
		}
	}
	return count
}