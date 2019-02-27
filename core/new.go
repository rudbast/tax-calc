package core

import "errors"

type (
	TaxCategory interface {
		// Tax product's name.
		Name() string

		// Product refundable indicator.
		IsRefundable() bool

		// Product tax identifier.
		Code() TaxCode

		// Product tax identifier name.
		Type() string

		// Product price amount.
		Price() float64

		// Product tax amount.
		Tax() float64

		// Product total price amount including tax.
		Total() float64
	}

	TaxCode int
)

const (
	TaxCodeFoodNBeverage TaxCode = iota + 1
	TaxCodeTobacco
	TaxCodeEntertainment
)

var (
	ErrUnknownTaxCode = errors.New("core: unknown / invalid tax code")
)

// Create new tax category object based on given code.
func New(name string, code TaxCode, price float64) (TaxCategory, error) {
	var tax TaxCategory

	switch code {
	case TaxCodeFoodNBeverage:
		tax = &TaxCategoryFoodNBeverage{name, price}

	case TaxCodeTobacco:
		tax = &TaxCategoryTobacco{name, price}

	case TaxCodeEntertainment:
		tax = &TaxCategoryEntertainment{name, price}

	default:
		return nil, ErrUnknownTaxCode
	}

	return tax, nil
}
