package api

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	billService = &BillService{}
)

func TestBillServiceGetBillsSumary(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	billService.db = db

	rows := sqlmock.NewRows(
		[]string{"id", "name", "tax_code", "price"},
	).AddRow(
		int64(1), "foo", 1, 15.0,
	).AddRow(
		int64(2), "bar", 2, 0.55,
	)

	mock.ExpectQuery(".+").WillReturnRows(rows)

	summary, apiErr := billService.GetBillsSummary(context.Background())
	assert.Nil(t, apiErr)
	assert.Equal(t, 2, len(summary.Bills))
	assert.Equal(t, "15.55", summary.PriceSubtotal)
	assert.Equal(t, "11.51", summary.TaxSubtotal)

	require.NoError(t, mock.ExpectationsWereMet())
}
