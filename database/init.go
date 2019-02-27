package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Option struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

// Connect to database returning the instance.
func Connect(opt Option) (*sql.DB, error) {
	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Name)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
