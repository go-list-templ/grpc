package v1

import (
	"github.com/go-list-templ/grpc/internal/usecase"
	"github.com/go-list-templ/proto/gen/api/user/v1"
	"go.uber.org/zap"
)

type V1 struct {
	v1.UserServiceServer

	u usecase.User
	l zap.Logger
}
