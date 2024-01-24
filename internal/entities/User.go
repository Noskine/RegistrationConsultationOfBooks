package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"username"`
	Email    string             `bson:"email"`
	CrateAt  time.Time          `bson:"create_at"`
	Unique   bool               `bson:"unique"`
}

func NewUser(username, email string) *User {
	return &User{
		UserName: username,
		Email:    email,
		CrateAt:  time.Now(),
		Unique:   true,
	}
}
