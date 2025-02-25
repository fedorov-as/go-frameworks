package main

import (
	"chi_example/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	appHandler := handlers.NewAppHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", appHandler.Root)
	r.Post("/user", appHandler.AddUser)
	r.Post("/post", appHandler.AddPost)
	r.Get("/post/{id}", appHandler.GetPost)
	r.Get("/posts/{nickname}", appHandler.GetUserPosts)
	r.Delete("/post/{id}", appHandler.DeletePost)
	http.ListenAndServe(":3000", r)
}
