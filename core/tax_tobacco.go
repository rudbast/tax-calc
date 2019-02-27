package core

type (
	TaxCategoryTobacco struct {
		name  string
		price float64
	}
)

const (
	taxTypeTobacco = "Tobacco"
)

func (t *TaxCategoryTobacco) Name() string       { return t.name }
func (t *TaxCategoryTobacco) IsRefundable() bool { return true }
func (t *TaxCategoryTobacco) Code() TaxCode      { return TaxCodeTobacco }
func (t *TaxCategoryTobacco) Type() string       { return taxTypeTobacco }
func (t *TaxCategoryTobacco) Price() float64     { return t.price }
func (t *TaxCategoryTobacco) Tax() float64       { return 10 + (0.02 * t.price) }
func (t *TaxCategoryTobacco) Total() float64     { return t.price + t.Tax() }
