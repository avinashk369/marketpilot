package database

import (
	"context"
	"fmt"

	"github.com/fintech/marketpilot/packages/shared/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(
	ctx context.Context,
	cfg *config.Config,
) (*Database, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Database{
		Pool: pool,
	}, nil
}

func (d *Database) Close() {
	d.Pool.Close()
}