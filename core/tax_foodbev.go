package core

type (
	TaxCategoryFoodNBeverage struct {
		name  string
		price float64
	}
)

const (
	taxTypeFoodNBeverage = "Food & Beverage"
)

func (t *TaxCategoryFoodNBeverage) Name() string       { return t.name }
func (t *TaxCategoryFoodNBeverage) IsRefundable() bool { return false }
func (t *TaxCategoryFoodNBeverage) Code() TaxCode      { return TaxCodeFoodNBeverage }
func (t *TaxCategoryFoodNBeverage) Type() string       { return taxTypeFoodNBeverage }
func (t *TaxCategoryFoodNBeverage) Price() float64     { return t.price }
func (t *TaxCategoryFoodNBeverage) Tax() float64       { return 0.1 * t.price }
func (t *TaxCategoryFoodNBeverage) Total() float64     { return t.price + t.Tax() }
