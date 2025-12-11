package event

import "github.com/go-list-templ/grpc/internal/domain/vo"

const (
	UserCreated TypeUserEvent = "created"
	UserDeleted TypeUserEvent = "deleted"
)

type (
	TypeUserEvent string

	UserEvent struct {
		UserID    vo.ID
		EventTime vo.Time
		TypeUser  TypeUserEvent
	}
)
