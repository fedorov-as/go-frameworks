package handlers

import (
	"net/http"

	application "github.com/fedorov-as/app_example_logic/pkg/app"
	"github.com/labstack/echo/v4"
)

type AppHandler struct {
	app *application.UsersPosts
}

func NewAppHandler() *AppHandler {
	return &AppHandler{
		app: application.NewUsersPosts(),
	}
}

func (h AppHandler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "Main Page")
}

func (h AppHandler) AddUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	user, err := h.app.CreateUser(u.Nickname, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	u = &User{
		Nickname: user.Nickname,
	}

	return c.JSON(http.StatusCreated, u)
}

func (h AppHandler) AddPost(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(p); err != nil {
		return err
	}

	auth := c.Get("username").(string)
	if auth != p.OwnerNick {
		return echo.NewHTTPError(http.StatusForbidden, "owner mismatches auth")
	}

	post, err := h.app.AddPost(p.Text, p.OwnerNick)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	p = &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	return c.JSON(http.StatusCreated, p)
}

func (h AppHandler) GetPost(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	post, err := h.app.GetPost(p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	p = &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	return c.JSON(http.StatusOK, p)
}

func (h AppHandler) GetUserPosts(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	posts, err := h.app.GetUserPosts(u.Nickname)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := make([]Post, 0, len(posts))
	for _, p := range posts {
		response = append(response, Post{
			ID:        p.ID,
			Text:      p.Text,
			OwnerNick: p.Owner.Nickname,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h AppHandler) DeletePost(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.app.DeletePost(p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
