package vo

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	MinLength = 5
	MaxLength = 30
)

type Username struct {
	value string
}

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func NewUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)

	if len(username) < MinLength {
		return Username{}, fmt.Errorf("username must be at least %v characters", MinLength)
	}
	if len(username) > MaxLength {
		return Username{}, fmt.Errorf("username must be less than %v characters", MaxLength)
	}

	if !usernameRegex.MatchString(username) {
		return Username{}, errors.New("username can only contain letters, numbers and underscores")
	}

	return Username{value: username}, nil
}

func (u Username) Value() string {
	return u.value
}

func (u Username) Equals(other Username) bool {
	return u.value == other.value
}
