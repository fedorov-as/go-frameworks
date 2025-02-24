package repo

import (
	"example_logic/pkg/model"
)

type UsersRepo interface {
	AddUser(user model.User) error
	GetUser(nickname string) (model.User, error)
}
