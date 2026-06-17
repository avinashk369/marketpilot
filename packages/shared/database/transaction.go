package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (d *Database) BeginTx(
	ctx context.Context,
) (pgx.Tx, error) {

	return d.Pool.BeginTx(
		ctx,
		pgx.TxOptions{},
	)
}