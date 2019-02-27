package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCategoryFoodNBeverage(t *testing.T) {
	name := "Big Mac"
	price := float64(1000)

	product := &TaxCategoryFoodNBeverage{name, price}

	assert.Equal(t, name, product.Name())
	assert.Equal(t, true, product.IsRefundable())
	assert.Equal(t, TaxCodeFoodNBeverage, product.Code())
	assert.Equal(t, taxTypeFoodNBeverage, product.Type())
	assert.Equal(t, price, product.Price())
	assert.Equal(t, float64(100), product.Tax())
	assert.Equal(t, float64(1100), product.Total())
}
