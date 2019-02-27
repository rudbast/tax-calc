package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	product, err := New("", TaxCodeEntertainment, 150.0)
	assert.NoError(t, err)

	_, ok := product.(*TaxCategoryEntertainment)
	assert.True(t, ok)

	product, err = New("", TaxCodeFoodNBeverage, 150.0)
	assert.NoError(t, err)

	_, ok = product.(*TaxCategoryFoodNBeverage)
	assert.True(t, ok)

	product, err = New("", TaxCodeTobacco, 150.0)
	assert.NoError(t, err)

	_, ok = product.(*TaxCategoryTobacco)
	assert.True(t, ok)

	product, err = New("", TaxCode(-1), 150.0)
	assert.Equal(t, ErrUnknownTaxCode, err)
}
