package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rudbast/tax-calc/model"
	"github.com/rudbast/tax-calc/core"
)

type (
	BillService struct {
		db *sql.DB
	}

	BillDetail struct {
		Name       string `json:"name"`
		TaxCode    string `json:"tax_code"`
		Type       string `json:"type"`
		Refundable bool   `json:"refundable"`
		Price      string `json:"price"`
		Tax        string `json:"tax"`
		Amount     string `json:"amount"`
	}

	BillsSummary struct {
		Bills         []BillDetail `json:"bills"`
		PriceSubtotal string       `json:"price_subtotal"`
		TaxSubtotal   string       `json:"tax_subtotal"`
		GrandTotal    string       `json:"grand_total"`
	}
)

// Insert one bill detail to database.
func (b *BillService) InsertOneBill(ctx context.Context, bill model.BillModel) error {
	_, err := model.InsertOneBill(ctx, b.db, bill)
	if err != nil {
		return errors.Wrap(err, "service/bill/insert")
	}

	return nil
}

// Get summary of all bills from database.
func (b *BillService) GetBillsSummary(ctx context.Context) (BillsSummary, error) {
	bills, err := model.GetAllBills(ctx, b.db)
	if err != nil {
		return BillsSummary{}, errors.Wrap(err, "service/bill/summary")
	}

	var summary BillsSummary

	var priceSubtotal float64
	var taxSubtotal float64
	var grandTotal float64

	for _, b := range bills {
		bill, err := core.New(b.Name, core.TaxCode(b.TaxCode), b.Price)
		if err != nil {
			return BillsSummary{}, errors.Wrap(err, "service/bill/summary/new")
		}

		priceSubtotal += bill.Price()
		taxSubtotal += bill.Tax()
		grandTotal += bill.Total()

		summary.Bills = append(summary.Bills, BillDetail{
			Name:       bill.Name(),
			TaxCode:    fmt.Sprint(bill.Code()),
			Type:       bill.Type(),
			Refundable: bill.IsRefundable(),
			Price:      fmt.Sprintf("%.2f", bill.Price()),
			Tax:        fmt.Sprintf("%.2f", bill.Tax()),
			Amount:     fmt.Sprintf("%.2f", bill.Total()),
		})
	}

	summary.PriceSubtotal = fmt.Sprintf("%.2f", priceSubtotal)
	summary.TaxSubtotal = fmt.Sprintf("%.2f", taxSubtotal)
	summary.GrandTotal = fmt.Sprintf("%.2f", grandTotal)

	return summary, nil
}
