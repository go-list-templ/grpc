package v1

import (
	pbgrpc "google.golang.org/grpc"

	"github.com/go-list-templ/grpc/internal/usecase"
	"github.com/go-list-templ/proto/gen/api/user/v1"
	"go.uber.org/zap"
)

func NewUserRoutes(app *pbgrpc.Server, u usecase.User, l zap.Logger) {
	r := &V1{u: u, l: l}
	{
		v1.RegisterUserServiceServer(app, r)
	}
}
