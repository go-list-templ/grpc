package service

import (
	"context"

	v1 "github.com/go-list-templ/proto/gen/api/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) AllUsers(ctx context.Context, req *v1.AllUsersRequest) (*v1.AllUsersResponse, error) {
	return &v1.AllUsersResponse{
		Users: []*v1.User{
			{
				Id:       1,
				Username: "test",
				Email:    "test@example.com",
			},
		},
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return &v1.CreateUserResponse{
		User: &v1.User{
			Id:       1,
			Username: req.Username,
			Email:    req.Email,
		},
	}, nil
}
