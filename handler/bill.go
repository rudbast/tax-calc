package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/rudbast/tax-calc/api"
	"github.com/rudbast/tax-calc/model"
)

type (
	BillModule struct {
		Service *api.BillService
	}
)

func NewBillModule(db *sql.DB) *BillModule {
	return &BillModule{
		Service: api.NewBillService(db),
	}
}

func (m BillModule) InsertBill(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	taxCode, err := strconv.Atoi(r.FormValue("tax_code"))
	if err != nil {
		return nil, api.NewErrorWrap(err, "param",
			"Invalid tax code value", http.StatusBadRequest)
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		return nil, api.NewErrorWrap(err, "param",
			"Invalid price value", http.StatusBadRequest)
	}

	bill := model.BillModel{
		Name:    r.FormValue("name"),
		TaxCode: taxCode,
		Price:   price,
	}

	apiErr := m.Service.InsertOneBill(ctx, bill)
	if apiErr != nil {
		return nil, apiErr
	}

	return nil, nil
}

func (m BillModule) GetBillsSummary(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	summary, apiErr := m.Service.GetBillsSummary(ctx)
	if apiErr != nil {
		return nil, apiErr
	}

	return summary, nil
}
