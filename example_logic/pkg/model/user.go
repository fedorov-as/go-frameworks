package model

type User struct {
	Nickname string
	password string
}

func NewUser(nickname, password string) User {
	return User{
		Nickname: nickname,
		password: password,
	}
}
