package vo

import (
	"errors"

	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func NewIDFromString(id string) (ID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return ID{}, errors.New("invalid id format")
	}

	return ID{value: parsed}, nil
}

func (id ID) Value() uuid.UUID {
	return id.value
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
}
