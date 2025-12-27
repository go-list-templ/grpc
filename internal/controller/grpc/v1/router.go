package v1

import (
	v1 "github.com/go-list-templ/proto/gen/api/user/v1"
	pbgrpc "google.golang.org/grpc"

	"github.com/go-list-templ/grpc/internal/usecase"
	"go.uber.org/zap"
)

func NewUserRoutes(app *pbgrpc.Server, u usecase.User, l zap.Logger) {
	r := &V1{u: u, l: l}
	{
		v1.RegisterUserServiceServer(app, r)
	}
}
