package v1

import (
	v1 "github.com/go-list-templ/proto/gen/api/user/v1"

	"github.com/go-list-templ/grpc/internal/usecase"
	"go.uber.org/zap"
)

type V1 struct {
	v1.UserServiceServer

	u usecase.User
	l zap.Logger
}
