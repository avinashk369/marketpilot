package server

import (
	"context"
	"net/http"
	"time"

	"github.com/fintech/marketpilot/apps/api/internal/bootstrap"
)

type Server struct {
	container  *bootstrap.Container
	httpServer *http.Server
}

func New(container *bootstrap.Container) *Server {

	router := setupRouter(container)

	return &Server{
		container: container,
		httpServer: &http.Server{
			Addr:              ":8080",
			Handler:           router,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}