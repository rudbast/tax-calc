package model

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type BillModel struct {
	ID      int64
	Name    string
	TaxCode int
	Price   float64
}

// Query all bill from database.
func GetAllBills(ctx context.Context, db *sql.DB) ([]BillModel, error) {
	rows, err := db.QueryContext(ctx, queryGetAllBills)
	if err != nil {
		return nil, errors.Wrap(err, "model/bill/all")
	}
	defer rows.Close()

	var bills []BillModel

	for rows.Next() {
		var bill BillModel

		err = rows.Scan(&bill.ID, &bill.Name, &bill.TaxCode, &bill.Price)
		if err != nil {
			return bills, errors.Wrap(err, "model/bill/all")
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

// Insert one bill to database, returning the modified model with inserted row identifier.
func InsertOneBill(ctx context.Context, db *sql.DB, bill BillModel) (BillModel, error) {
	query := queryInsertOneBill + " RETURNING id"

	err := db.QueryRowContext(ctx, query, bill.Name, bill.TaxCode, bill.Price).Scan(&bill.ID)
	if err != nil {
		return bill, errors.Wrap(err, "model/bill/insert")
	}

	return bill, nil
}
