package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	application "github.com/fedorov-as/app_example_logic/pkg/app"
	"github.com/go-chi/chi/v5"
)

type AppHandler struct {
	app *application.UsersPosts
}

func NewAppHandler() *AppHandler {
	return &AppHandler{
		app: application.NewUsersPosts(),
	}
}

func (h AppHandler) sendResponse(w http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h AppHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Main Page"))
}

func (h AppHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := h.app.CreateUser(u.Nickname, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	u = &User{
		Nickname: user.Nickname,
	}

	h.sendResponse(w, u)
}

func (h AppHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	p := new(Post)
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	post, err := h.app.AddPost(p.Text, p.OwnerNick)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	p = &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	h.sendResponse(w, p)
}

func (h AppHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	post, err := h.app.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	p := &Post{
		ID:        post.ID,
		Text:      post.Text,
		OwnerNick: post.Owner.Nickname,
	}

	h.sendResponse(w, p)
}

func (h AppHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "nickname")

	posts, err := h.app.GetUserPosts(owner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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

	h.sendResponse(w, response)
}

func (h AppHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.app.DeletePost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	h.sendResponse(w, "delete OK")
}
