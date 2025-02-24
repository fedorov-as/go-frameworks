package repo

import (
	"example_logic/pkg/model"
	"fmt"
	"slices"
)

type UsersMemoryRepo struct {
	data []model.User
}

var _ UsersRepo = &UsersMemoryRepo{}

func NewUsersMemoryRepo() *UsersMemoryRepo {
	return &UsersMemoryRepo{
		data: make([]model.User, 0),
	}
}

func (repo *UsersMemoryRepo) AddUser(user model.User) error {
	if slices.IndexFunc(repo.data, func(elem model.User) bool { return elem.Nickname == user.Nickname }) != -1 {
		return fmt.Errorf("user exists")
	}

	repo.data = append(repo.data, user)
	return nil
}

func (repo UsersMemoryRepo) GetUser(nickname string) (model.User, error) {
	ind := slices.IndexFunc(repo.data, func(elem model.User) bool { return elem.Nickname == nickname })

	if ind == -1 {
		return model.User{}, fmt.Errorf("no user")
	}

	return repo.data[ind], nil
}
