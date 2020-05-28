package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{Name: "Nic", Price: 1.00, SKU: "abs-abc-def"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
