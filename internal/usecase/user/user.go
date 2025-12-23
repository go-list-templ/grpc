package user

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
	"github.com/go-list-templ/grpc/internal/repo"
)

type UseCase struct {
	repo repo.UserPersistentRepo
}

func New(repo repo.UserPersistentRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) All(ctx context.Context) ([]entity.User, error) {
	users, err := u.repo.All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	createdUser, err := u.repo.Store(ctx, user)
	if err != nil {
		return user, err
	}

	return createdUser, nil
}

func (u *UseCase) Delete(ctx context.Context, userID vo.ID) error {
	return u.repo.Destroy(ctx, userID)
}

func (u *UseCase) Show(ctx context.Context, userID vo.ID) (entity.User, error) {
	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UseCase) Update(ctx context.Context, user entity.User) (entity.User, error) {
	updatedUser, err := u.repo.Change(ctx, user)
	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
