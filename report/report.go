package report

type Valuable interface {
	Total() int
}

func Sum(items []Valuable) int {
	total := 0
	for i := 0; i < len(items); i++ {
		total += items[i].Total()
	}
	return total
}

func CountAbove(items []Valuable, threshold int) int {
	count := 0
	for i := 0; i < len(items); i++ {
		if items[i].Total() > threshold {
			count++
		}
	}
	return count
}
