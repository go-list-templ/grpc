package persistent

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
	"github.com/go-list-templ/grpc/internal/infra/persistent/postgres"
)

type UserPostgresRepo struct {
	*postgres.Postgres
}

func NewUserPostgresRepo(postgres *postgres.Postgres) *UserPostgresRepo {
	return &UserPostgresRepo{postgres}
}

func (r *UserPostgresRepo) Store(ctx context.Context) (entity.User, error) {
	return entity.User{}, nil
}

func (r *UserPostgresRepo) Change(ctx context.Context, user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (r *UserPostgresRepo) Destroy(ctx context.Context, user entity.User) error {
	return nil
}

func (r *UserPostgresRepo) GetById(ctx context.Context, id vo.ID) (entity.User, error) {
	return entity.User{}, nil
}

func (r *UserPostgresRepo) All(ctx context.Context) ([]entity.User, error) {
	return []entity.User{}, nil
}
