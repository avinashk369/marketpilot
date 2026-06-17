package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	instance *zap.Logger
	once     sync.Once
)

func New() (*zap.Logger, error) {
	var err error

	once.Do(func() {
		instance, err = zap.NewProduction()
	})

	return instance, err
}