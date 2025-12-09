package entity

import "github.com/go-list-templ/grpc/internal/domain/vo"

type User struct {
	ID        vo.ID
	Username  vo.Username
	Email     vo.Email
	CreatedAt vo.Time
	UpdatedAt vo.Time
}

func NewUser(username, email string) (*User, error) {
	validUsername, err := vo.NewUsername(username)
	if err != nil {
		return nil, err
	}

	validEmail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        vo.NewID(),
		Username:  validUsername,
		Email:     validEmail,
		CreatedAt: vo.NewTime(),
		UpdatedAt: vo.NewTime(),
	}, nil
}
