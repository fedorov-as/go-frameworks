package handlers

type User struct {
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}
