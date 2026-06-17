package bootstrap

import (
	"context"

	"github.com/fintech/marketpilot/packages/shared/cache"
	"github.com/fintech/marketpilot/packages/shared/config"
	"github.com/fintech/marketpilot/packages/shared/database"
	"github.com/fintech/marketpilot/packages/shared/logger"
	"go.uber.org/zap"
)

type Container struct {
	Config   *config.Config
	Logger   *zap.Logger
	Database *database.Database
	Cache    *cache.Cache
}

func NewContainer() (*Container, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	log, err := logger.New()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	db, err := database.New(
		ctx,
		cfg,
	)
	if err != nil {
		return nil, err
	}

	redisCache, err := cache.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Container{
		Config:   cfg,
		Logger:   log,
		Database: db,
		Cache:    redisCache,
	}, nil
}