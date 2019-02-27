package model

var (
	queryGetAllBills string = `
		SELECT
			id,
			name,
			tax_code,
			price
		FROM
			bills`

	queryInsertOneBill string = `
		INSERT INTO bills (
			name, tax_code, price
		) VALUES (
			$1, $2, $3
		)`
)
