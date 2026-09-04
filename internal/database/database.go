package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
	Ctx  context.Context
}

func NewDB(ctx context.Context) (*DB, error) {

	conn_string := os.Getenv("CONN_STRING")
	pool, err := pgxpool.New(ctx, conn_string)
	if err != nil {
		return nil, err
	}

	return &DB{
		Pool: pool,
		Ctx:  ctx,
	}, nil

}

func (db *DB) Close() error {

	if db.Pool != nil {
		db.Pool.Close()
	}
	return nil

}
