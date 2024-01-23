package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"username"`
	Email    string             `bson:"email"`
	CrateAt  primitive.DateTime `bson:"create_at"`
}

func NewUser(username, email string) *User {
	return &User{
		UserName: username,
		Email:    email,
	}
}
