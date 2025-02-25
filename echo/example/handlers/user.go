package handlers

type User struct {
	Nickname string `json:"nickname" param:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}
