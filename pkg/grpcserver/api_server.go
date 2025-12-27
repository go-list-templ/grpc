package grpcserver

import (
	"net"

	"github.com/go-list-templ/grpc/config"
	"google.golang.org/grpc"
)

type APIServer struct {
	server *grpc.Server
	config *config.Server
	errors chan error
}

func NewAPIServer(cfg *config.Server) *APIServer {
	grpcServer := grpc.NewServer()

	return &APIServer{
		server: grpcServer,
		config: cfg,
		errors: make(chan error, 1),
	}
}

func (s *APIServer) Notify() <-chan error {
	return s.errors
}

func (s *APIServer) Start() {
	go func() {
		lis, err := net.Listen("tcp", net.JoinHostPort("", s.config.GRPCPort))
		if err != nil {
			s.errors <- err
		}

		s.errors <- s.server.Serve(lis)
		close(s.errors)
	}()
}

func (s *APIServer) Stop() {
	s.server.GracefulStop()
}
