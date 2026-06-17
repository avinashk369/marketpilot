package cache

import "context"

func (c *Cache) Health(
	ctx context.Context,
) error {
	return c.Client.Ping(ctx).Err()
}