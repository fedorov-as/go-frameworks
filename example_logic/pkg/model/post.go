package model

type Post struct {
	ID    int
	Text  string
	Owner User
}

func NewPost(id int, text string, owner User) Post {
	return Post{
		ID:    id,
		Text:  text,
		Owner: owner,
	}
}
