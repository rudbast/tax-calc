package model

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllBills(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{"id", "name", "tax_code", "price"},
	).AddRow(
		int64(1), "foo", 1, 15.0,
	).AddRow(
		int64(2), "bar", 2, 0.55,
	)

	mock.ExpectQuery("SELECT (.+) FROM bills").WillReturnRows(rows)

	bills, err := GetAllBills(context.Background(), db)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(bills))

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertOneBill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bill := BillModel{0, "foo", 1, 15.0}

	rows := sqlmock.NewRows(
		[]string{"id"},
	).AddRow(int64(1))

	mock.ExpectQuery("INSERT INTO bills (.+) RETURNING id").WithArgs(
		bill.Name, bill.TaxCode, bill.Price,
	).WillReturnRows(rows)

	insertedBill, err := InsertOneBill(context.Background(), db, bill)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), insertedBill.ID)

	require.NoError(t, mock.ExpectationsWereMet())
}
