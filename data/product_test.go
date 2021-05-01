package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "mocha latte",
		Price: 1.19,
		SKU:   "asd-dfg-ert",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
