package rental

import "errors"

type Item struct {
	Code      string
	Name      string
	DailyRate int
	Available int
	Rented    int
}

func (i Item) Value() int {
	return i.DailyRate * i.Available
}

func (i *Item) Rent(units int) error {
	if units <= 0 {
		return errors.New("amount must be positive")
	} else if units > i.Available {
		return errors.New("insufficient equipment")
	} else {
		i.Available -= units
		i.Rented += units
		return nil
	}
}

func (i *Item) Return(units int) error {
	if units <= 0 {
		return errors.New("amount must be positive")
	} else if units > i.Rented {
		return errors.New("insufficient rented units")
	} else {
		i.Available += units
		i.Rented -= units
		return nil
	}
}
