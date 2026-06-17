package database

import "context"

func (d *Database) Health(
	ctx context.Context,
) error {
	return d.Pool.Ping(ctx)
}