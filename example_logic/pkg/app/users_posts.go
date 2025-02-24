package app

import (
	"example_logic/pkg/model"
	"example_logic/pkg/repo"
)

type UsersPosts struct {
	users repo.UsersRepo
	posts repo.PostsRepo
}

func NewUsersPosts() *UsersPosts {
	return &UsersPosts{
		users: repo.NewUsersMemoryRepo(),
		posts: repo.NewPostsMemoryRepo(),
	}
}

func (app *UsersPosts) CreateUser(nickname, password string) error {
	return app.users.AddUser(model.NewUser(nickname, password))
}

func (app *UsersPosts) AddPost(text string, nickname string) error {
	owner, err := app.users.GetUser(nickname)
	if err != nil {
		return err
	}

	return app.posts.AddPost(text, owner)
}

func (app UsersPosts) GetPost(id int) (model.Post, error) {
	return app.posts.GetPostByID(id)
}

func (app UsersPosts) GetUserPosts(nickname string) ([]model.Post, error) {
	owner, err := app.users.GetUser(nickname)
	if err != nil {
		return nil, err
	}

	return app.posts.GetPostsByOwner(owner)
}
