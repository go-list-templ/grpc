package usecase

import "github.com/go-list-templ/grpc/internal/domain/entity"

type (
	User interface {
		Create() (entity.User, error)
		Update(user entity.User) (entity.User, error)
		Delete(user entity.User) error
		Show() (entity.User, error)
		All() ([]entity.User, error)
	}
)
