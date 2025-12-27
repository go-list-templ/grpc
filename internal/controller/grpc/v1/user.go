package v1

import (
	"context"
	"fmt"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	v1 "github.com/go-list-templ/proto/gen/api/user/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *V1) CreateUser(ctx context.Context, request *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user, err := entity.NewUser(request.GetUsername(), request.GetEmail())
	if err != nil {
		r.l.Error("grpc - v1 - NewUser", zap.Error(err))

		return nil, fmt.Errorf("grpc - v1 - NewUser: %w", err)
	}

	createdUser, err := r.u.Create(ctx, *user)
	if err != nil {
		r.l.Error("grpc - v1 - CreateUser", zap.Error(err))

		return nil, fmt.Errorf("grpc - v1 - CreateUser: %w", err)
	}

	return &v1.CreateUserResponse{
		User: &v1.User{
			Id:        createdUser.ID.Value().ID(),
			Username:  createdUser.Name.Value(),
			Email:     createdUser.Email.Value(),
			CreatedAt: timestamppb.New(createdUser.CreatedAt),
			UpdatedAt: timestamppb.New(createdUser.UpdatedAt),
		},
	}, nil
}
