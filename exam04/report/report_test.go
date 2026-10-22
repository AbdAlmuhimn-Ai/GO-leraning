package report

import (
	"testing"

	"example.com/academy/exam04/rental"
)

func TestSum(t *testing.T) {
	items := []Valuable{
		rental.Item{DailyRate: 30, Available: 3},
		rental.Item{DailyRate: 10, Available: 2},
		rental.Service{Fee: 15, Sessions: 2},
	}
	empty := []Valuable{}

	if Sum(items) != 140 || Sum(empty) != 0 {
		t.Fatalf("Sum failed for items or empty slice")
	}
}

func TestCountAbove(t *testing.T) {
	items := []Valuable{
		rental.Item{DailyRate: 30, Available: 3},
		rental.Item{DailyRate: 10, Available: 2},
		rental.Service{Fee: 15, Sessions: 2},
	}
	empty := []Valuable{}

	if CountAbove(items, 25) != 2 || CountAbove(empty, 25) != 0 {
		t.Fatalf("CountAbove failed for threshold 25 or empty slice")
	}
}