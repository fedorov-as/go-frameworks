package handlers

import (
	"net/http"

	application "github.com/fedorov-as/app_example_logic/pkg/app"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	app *application.UsersPosts
}

func NewAppHandler() *AppHandler {
	return &AppHandler{
		app: application.NewUsersPosts(),
	}
}

func (h AppHandler) Root(c *gin.Context) {
	c.JSON(http.StatusOK, "Main Page")
}

func (h AppHandler) AddUser(c *gin.Context) {
	u := new(User)
	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.app.CreateUser(u.Nickname, u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	u = &User{
		Nickname: user.Nickname,
	}

	c.JSON(http.StatusCreated, u)
}

func (h AppHandler) AddPost(c *gin.Context) {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.app.AddPost(p.Text, p.OwnerNick)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	p = &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	c.JSON(http.StatusCreated, p)
}

func (h AppHandler) GetPost(c *gin.Context) {
	p := new(Post)
	if err := c.BindUri(p); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.app.GetPost(p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	p = &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	c.JSON(http.StatusOK, p)
}

func (h AppHandler) GetUserPosts(c *gin.Context) {
	u := new(User)
	if err := c.BindUri(u); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	posts, err := h.app.GetUserPosts(u.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := make([]Post, 0, len(posts))
	for _, p := range posts {
		response = append(response, Post{
			ID:        p.ID,
			Text:      p.Text,
			OwnerNick: p.Owner.Nickname,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h AppHandler) DeletePost(c *gin.Context) {
	p := new(Post)
	if err := c.BindUri(p); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.app.DeletePost(p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
