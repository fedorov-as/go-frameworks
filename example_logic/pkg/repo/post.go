package repo

import "example_logic/pkg/model"

type PostsRepo interface {
	AddPost(text string, owner model.User) error
	GetPostByID(id int) (model.Post, error)
	GetPostsByOwner(owner model.User) ([]model.Post, error)
}
