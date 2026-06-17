package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fintech/marketpilot/apps/api/internal/bootstrap"
	"github.com/fintech/marketpilot/apps/api/internal/server"
)

func main() {

	container, err := bootstrap.NewContainer()
	if err != nil {
		panic(err)
	}

	srv := server.New(container)

	go func() {

		container.Logger.Info("api service starting")

		if err := srv.Start(); err != nil {
			container.Logger.Error(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	container.Logger.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		container.Logger.Error(err.Error())
	}
}