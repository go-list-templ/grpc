package grpc

import (
	"net"

	"github.com/go-list-templ/grpc/config"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	config *config.Server
	errors chan error
}

func NewServer(cfg *config.Server) *Server {
	grpcServer := grpc.NewServer()

	return &Server{
		server: grpcServer,
		config: cfg,
		errors: make(chan error, 1),
	}
}

func (s *Server) Notify() <-chan error {
	return s.errors
}

func (s *Server) Start() {
	go func() {
		lis, err := net.Listen("tcp", net.JoinHostPort("", s.config.GRPCPort))
		if err != nil {
			s.errors <- err
		}

		s.errors <- s.server.Serve(lis)
		close(s.errors)
	}()
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
