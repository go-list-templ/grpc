package repo

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
)

type (
	UserPersistentRepo interface {
		Store(context.Context) (entity.User, error)
		Change(context.Context, entity.User) (entity.User, error)
		Destroy(context.Context, entity.User) error
		GetById(context.Context, vo.ID) (entity.User, error)
		All(context.Context) ([]entity.User, error)
	}

	UserExternalRepo interface{}
)
