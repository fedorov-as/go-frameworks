package main

import (
	"echo_example/handlers"
	"echo_example/middlewares"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	appHandler := handlers.NewAppHandler()

	e := echo.New()
	e.Validator = &handlers.CustomValidator{Validator: validator.New()}

	// e.Use(middlewares.AuthMiddleware)

	e.GET("/", appHandler.Root)
	e.POST("/user", appHandler.AddUser)
	e.POST("/post", appHandler.AddPost, middlewares.AuthMiddleware)
	e.GET("/post/:id", appHandler.GetPost)
	e.GET("/posts/:nickname", appHandler.GetUserPosts)
	e.DELETE("/post/:id", appHandler.DeletePost)
	e.Logger.Fatal(e.Start(":3000"))
}
