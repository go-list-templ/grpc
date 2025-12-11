package event

import "github.com/go-list-templ/grpc/internal/domain/vo"

type TypeUserEvent string

const (
	UserCreated TypeUserEvent = "created"
	UserDeleted TypeUserEvent = "deleted"
)

type UserEvent struct {
	UserID    vo.ID
	EventTime vo.Time
	TypeUser  TypeUserEvent
}
