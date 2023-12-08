package entities

import "github.com/google/uuid"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Function string `json:"function"`
}

func NewUser(username, email, function, password_hash string) *User {
	return &User{
		Id:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: password_hash,
		Function: function,
	}
}

