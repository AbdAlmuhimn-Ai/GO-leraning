package rental

import "testing"

func TestRent(t *testing.T) {
	item := Item{Code: "CAM", DailyRate: 30, Available: 4, Rented: 0}

	if item.Rent(2) != nil || item.Available != 2 || item.Rented != 2 {
		t.Fatalf("Rent(2) failed to update fields correctly")
	}

	if item.Rent(0) == nil || item.Rent(-1) == nil || item.Rent(5) == nil || item.Available != 2 || item.Rented != 2 {
		t.Fatalf("Rent should reject 0, negative, and excess without changing data")
	}
}

func TestReturn(t *testing.T) {
	item := Item{Code: "CAM", DailyRate: 30, Available: 2, Rented: 2}

	if item.Return(1) != nil || item.Available != 3 || item.Rented != 1 {
		t.Fatalf("Return(1) failed to update fields correctly")
	}

	if item.Return(0) == nil || item.Return(-1) == nil || item.Return(5) == nil || item.Available != 3 || item.Rented != 1 {
		t.Fatalf("Return should reject invalid inputs without corrupting data")
	}
}

func TestFindIndex(t *testing.T) {
	items := []Item{
		{Code: "CAM"},
		{Code: "MIC"},
	}

	if FindIndex(items, "CAM") != 0 || FindIndex(items, "MIC") != 1 || FindIndex(items, "BAD") != -1 {
		t.Fatalf("FindIndex returned incorrect index")
	}
}

func TestRentByCode(t *testing.T) {
	items := []Item{
		{Code: "CAM", Available: 4, Rented: 0},
	}

	if RentByCode(items, "CAM", 2) != nil || items[0].Available != 2 || items[0].Rented != 2 {
		t.Fatalf("RentByCode failed on existing item")
	}

	if RentByCode(items, "BAD", 1) == nil || items[0].Available != 2 || items[0].Rented != 2 {
		t.Fatalf("RentByCode failed on non-existing item")
	}
}

func TestReturnByCode(t *testing.T) {
	items := []Item{
		{Code: "CAM", Available: 2, Rented: 2},
	}

	if ReturnByCode(items, "CAM", 1) != nil || items[0].Available != 3 || items[0].Rented != 1 {
		t.Fatalf("ReturnByCode failed on existing item")
	}

	if ReturnByCode(items, "BAD", 1) == nil || items[0].Available != 3 || items[0].Rented != 1 {
		t.Fatalf("ReturnByCode failed on non-existing item")
	}
}
