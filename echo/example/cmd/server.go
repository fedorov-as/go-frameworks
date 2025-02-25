package main

import (
	"echo_example/handlers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	appHandler := handlers.NewAppHandler()

	e := echo.New()
	e.Validator = &handlers.CustomValidator{Validator: validator.New()}
	e.GET("/", appHandler.Root)
	e.POST("/user", appHandler.AddUser)
	e.POST("/post", appHandler.AddPost)
	e.GET("/post/:id", appHandler.GetPost)
	e.GET("/posts/:nickname", appHandler.GetUserPosts)
	e.DELETE("/post/:id", appHandler.DeletePost)
	e.Logger.Fatal(e.Start(":1323"))
}
