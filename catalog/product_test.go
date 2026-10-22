package catalog

import "testing"

func TestTotal(t *testing.T) {
	p := Product{Price: 20, Quantity: 3}

	got := p.Total()
	want := 60

	if got != want {
		t.Fatalf("Total() = %d, want %d", got, want)
	}
}

func TestSell(t *testing.T) {
	p := Product{Price: 20, Quantity: 3}
	err := p.Sell(2)
	if err != nil {
		t.Fatalf("Sell() = %v, want nil", err)
	}
	if p.Quantity != 1 {
		t.Fatalf("Sell() = %d, want 1", p.Quantity)
	}

	err = p.Sell(0)
	if err == nil {
		t.Fatalf("Sell(0) = nil, want error")
	}
	if p.Quantity != 1 {
		t.Fatalf("Quantity changed after failed Sell(0): got %d, want 1", p.Quantity)
	}

	err = p.Sell(-1)
	if err == nil {
		t.Fatalf("Sell(-1) = nil, want error")
	}
	if p.Quantity != 1 {
		t.Fatalf("Quantity changed after failed Sell(-1): got %d, want 1", p.Quantity)
	}

	err = p.Sell(5)
	if err == nil {
		t.Fatalf("Sell(5) = nil, want error")
	}
	if p.Quantity != 1 {
		t.Fatalf("Quantity changed after failed Sell(5): got %d, want 1", p.Quantity)
	}
}

func TestRestock(t *testing.T) {
	p := Product{Price: 20, Quantity: 3}

	err := p.Restock(0)
	if err == nil {
		t.Fatalf("Restock(0) = nil, want error")
	}
	if p.Quantity != 3 {
		t.Fatalf("Quantity changed after failed Restock(0): got %d, want 3", p.Quantity)
	}

	err = p.Restock(-1)
	if err == nil {
		t.Fatalf("Restock(-1) = nil, want error")
	}
	if p.Quantity != 3 {
		t.Fatalf("Quantity changed after failed Restock(-1): got %d, want 3", p.Quantity)
	}

	err = p.Restock(4)
	if err != nil {
		t.Fatalf("Restock() = %v, want nil", err)
	}
	if p.Quantity != 7 {
		t.Fatalf("Restock() = %d, want 7", p.Quantity)
	}
}
