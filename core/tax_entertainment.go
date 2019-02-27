package core

type (
	TaxCategoryEntertainment struct {
		name  string
		price float64
	}
)

const (
	taxTypeEntertainment = "Entertainment"
)

func (t *TaxCategoryEntertainment) Name() string       { return t.name }
func (t *TaxCategoryEntertainment) IsRefundable() bool { return false }
func (t *TaxCategoryEntertainment) Code() TaxCode      { return TaxCodeEntertainment }
func (t *TaxCategoryEntertainment) Type() string       { return taxTypeEntertainment }
func (t *TaxCategoryEntertainment) Price() float64     { return t.price }
func (t *TaxCategoryEntertainment) Total() float64     { return t.price + t.Tax() }

func (t *TaxCategoryEntertainment) Tax() float64 {
	if t.price > 0 && t.price < 100 {
		return 0
	} else if t.price >= 100 {
		return 0.01 + (t.price - 100)
	}
	return 0
}
