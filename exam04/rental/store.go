package rental

import "errors"

func FindIndex(items []Item, code string) int {
	for i := 0; i < len(items); i++ {
		if items[i].Code == code {
			return i
		}
	}
	return -1
}

func RentByCode(items []Item, code string, units int) error {
	idx := FindIndex(items, code)
	if idx == -1 {
		return errors.New("item not found")
	}
	return items[idx].Rent(units)
}

func ReturnByCode(items []Item, code string, units int) error {
	idx := FindIndex(items, code)
	if idx == -1 {
		return errors.New("item not found")
	}
	return items[idx].Return(units)
}
