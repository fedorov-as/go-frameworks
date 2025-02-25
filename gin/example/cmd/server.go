package main

import (
	"gin_example/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	appHandler := handlers.NewAppHandler()

	r := gin.Default()
	r.GET("/", appHandler.Root)
	r.POST("/user", appHandler.AddUser)
	r.POST("/post", appHandler.AddPost)
	r.GET("/post/:id", appHandler.GetPost)
	r.GET("/posts/:nickname", appHandler.GetUserPosts)
	r.DELETE("/post/:id", appHandler.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
