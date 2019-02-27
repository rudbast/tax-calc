package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rudbast/tax-calc/core"
	"github.com/rudbast/tax-calc/model"
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

func NewBillService(db *sql.DB) *BillService {
	return &BillService{
		db: db,
	}
}

// Insert one bill detail to database.
func (b *BillService) InsertOneBill(ctx context.Context, bill model.BillModel) *Error {
	_, err := model.InsertOneBill(ctx, b.db, bill)
	if err != nil {
		return NewErrorWrap(err, "service/bill/insert",
			"Insert bill error.", http.StatusInternalServerError)
	}

	return nil
}

// Get summary of all bills from database.
func (b *BillService) GetBillsSummary(ctx context.Context) (BillsSummary, *Error) {
	bills, err := model.GetAllBills(ctx, b.db)
	if err != nil {
		return BillsSummary{}, NewErrorWrap(err, "service/bill/summary",
			"Query bill data error.", http.StatusInternalServerError)
	}

	var summary BillsSummary

	var priceSubtotal float64
	var taxSubtotal float64
	var grandTotal float64

	for _, b := range bills {
		bill, err := core.New(b.Name, core.TaxCode(b.TaxCode), b.Price)
		if err != nil {
			return BillsSummary{}, NewErrorWrap(err, "service/bill/summary/new",
				"Process bill data error.", http.StatusInternalServerError)
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
