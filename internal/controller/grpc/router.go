package grpc

import (
	pbgrpc "google.golang.org/grpc"

	"github.com/go-list-templ/grpc/internal/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
)

func NewRouter(app *pbgrpc.Server, u usecase.User, l zap.Logger) {
	{
		v1.NewTranslationRoutes(app, u, l)
	}

	reflection.Register(app)
}
