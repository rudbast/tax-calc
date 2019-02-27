package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCategoryEntertainment(t *testing.T) {
	name := "Movie"
	price := float64(150)

	product := &TaxCategoryEntertainment{name, price}

	assert.Equal(t, name, product.Name())
	assert.Equal(t, false, product.IsRefundable())
	assert.Equal(t, TaxCodeEntertainment, product.Code())
	assert.Equal(t, taxTypeEntertainment, product.Type())
	assert.Equal(t, price, product.Price())
	assert.Equal(t, 0.5, product.Tax())
	assert.Equal(t, 150.5, product.Total())

	product.price = float64(99)

	assert.Equal(t, float64(0), product.Tax())
}
