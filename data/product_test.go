package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 0.99,
		SKU:   "asb-asd-123",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidPriceReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: -1,
		SKU:   "asb-asd-123",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 0.99,
		SKU:   "asb",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		{
			Name:  "abc",
			Price: 0.99,
			SKU:   "asb-asd-123",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}
