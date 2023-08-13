package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgresPool struct {
	*pgx.Conn
}

func NewPostgressPool() (*PostgresPool, error) {
	ctx := context.Background()

	pgdb, err := pgx.Connect(ctx, os.Getenv("PGX_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgresPool{pgdb}, nil
}
