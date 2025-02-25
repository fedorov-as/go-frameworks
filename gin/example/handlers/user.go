package handlers

type User struct {
	Nickname string `json:"nickname" uri:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}
