package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCategoryTobacco(t *testing.T) {
	name := "Lucky Stretch"
	price := float64(1000)

	product := &TaxCategoryTobacco{name, price}

	assert.Equal(t, name, product.Name())
	assert.Equal(t, false, product.IsRefundable())
	assert.Equal(t, TaxCodeTobacco, product.Code())
	assert.Equal(t, taxTypeTobacco, product.Type())
	assert.Equal(t, price, product.Price(), price)
	assert.Equal(t, float64(30), product.Tax())
	assert.Equal(t, float64(1030), product.Total())
}
