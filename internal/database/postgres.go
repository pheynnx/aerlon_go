package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLXPool struct {
	*sqlx.DB
}

func NewSQLXPool() (*SQLXPool, error) {

	db, err := sqlx.Connect("postgres", os.Getenv("SQL_URL"))
	if err != nil {
		return nil, err
	}

	return &SQLXPool{db}, nil
}
