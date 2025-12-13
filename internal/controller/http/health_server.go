package http

import (
	"context"
	"net"
	"net/http"

	"github.com/go-list-templ/grpc/config"
)

type HealthServer struct {
	server http.Server
	config *config.Server
	errors chan error
}

func NewHealthServer(cfg *config.Server) *HealthServer {
	return &HealthServer{
		server: http.Server{
			Addr:              net.JoinHostPort("", cfg.HealthPort),
			Handler:           nil,
			ReadHeaderTimeout: cfg.HTTPTimeout,
			IdleTimeout:       cfg.IdleTimeout,
		},
		config: cfg,
		errors: make(chan error, 1),
	}
}

func (s *HealthServer) Notify() <-chan error {
	return s.errors
}

func (s *HealthServer) Start() {
	go func() {
		s.errors <- s.server.ListenAndServe()
		close(s.errors)
	}()
}

func (s *HealthServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
