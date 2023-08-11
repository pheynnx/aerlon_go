package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresPool struct {
	pgdb    *pgx.Conn
	pgdbCtx context.Context
}

func NewPostgressPool(connString string) (*PostgresPool, error) {
	pgxdbCtx := context.Background()

	pgdb, err := pgx.Connect(pgxdbCtx, connString)
	if err != nil {
		return nil, err
	}

	return &PostgresPool{
		pgdb, pgxdbCtx,
	}, nil

}
