package user

import "github.com/go-list-templ/grpc/internal/repo"

type UseCase struct {
	repo repo.UserPersistentRepo
}

func New(repo repo.UserPersistentRepo) *UseCase {
	return &UseCase{repo: repo}
}
