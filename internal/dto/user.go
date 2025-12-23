package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserInput struct {
	ID uuid.UUID `json:"id"`
}

type ShowUserInput struct {
	ID uuid.UUID `json:"id"`
}

type UserOutput struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
