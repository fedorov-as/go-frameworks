package repo

import (
	"example_logic/pkg/model"
	"fmt"
)

type PostsMemoryRepo struct {
	data   map[int]model.Post
	nextID int
}

var _ PostsRepo = &PostsMemoryRepo{}

func NewPostsMemoryRepo() *PostsMemoryRepo {
	return &PostsMemoryRepo{
		data:   make(map[int]model.Post),
		nextID: 0,
	}
}

func (repo *PostsMemoryRepo) AddPost(text string, owner model.User) error {
	repo.data[repo.nextID] = model.NewPost(repo.nextID, text, owner)
	repo.nextID++
	return nil
}

func (repo PostsMemoryRepo) GetPostByID(id int) (model.Post, error) {
	if post, ok := repo.data[id]; ok {
		return post, nil
	}

	return model.Post{}, fmt.Errorf("no post with id %d", id)
}

func (repo PostsMemoryRepo) GetPostsByOwner(owner model.User) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	for _, post := range repo.data {
		posts = append(posts, post)
	}

	return posts, nil
}
