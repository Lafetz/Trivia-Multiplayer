package mock

import (
	"errors"

	"github.com/lafetz/trivia-go/internal/data"
)

var mockUser = &data.UserModel{
	ID:    1,
	Name:  "Alice",
	Email: "alice@example.com",
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return errors.New(`pq: duplicate key value violates unique constraint "users_email_key"`)
	default:
		return nil
	}
}
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@example.com":
		return 1, nil
	default:
		return 0, errors.New("auth err")
	}
}
func (m *UserModel) Get(id int) (*data.UserModel, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, errors.New("404")
	}
}
