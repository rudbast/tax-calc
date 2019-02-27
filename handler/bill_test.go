package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	billHandler *BillModule
)

func TestBillHandlerGetBillsSummary(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	billHandler = NewBillModule(db)

	rows := sqlmock.NewRows(
		[]string{"id", "name", "tax_code", "price"},
	).AddRow(
		int64(1), "foo", 1, 15.0,
	).AddRow(
		int64(2), "bar", 2, 0.55,
	)

	mock.ExpectQuery(".+").WillReturnRows(rows)

	req, err := http.NewRequest(http.MethodGet, "/bills", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := HandlerFunc(billHandler.GetBillsSummary)

	handler.ServeHTTP(recorder, req)

	expectedResponse := `{"data":{"bills":[{"name":"foo","tax_code":"1","type":"Food \u0026 Beverage","refundable":true,"price":"15.00","tax":"1.50","amount":"16.50"},{"name":"bar","tax_code":"2","type":"Tobacco","refundable":false,"price":"0.55","tax":"10.01","amount":"10.56"}],"price_subtotal":"15.55","tax_subtotal":"11.51","grand_total":"27.06"}}`

	assert.Equal(t, expectedResponse, strings.TrimSuffix(recorder.Body.String(), "\n"))
	assert.Equal(t, 200, recorder.Code)

	require.NoError(t, mock.ExpectationsWereMet())
}
