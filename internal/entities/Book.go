package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"book_name"`
	Author          string             `bson:"book_author"`
	Category        string             `bson:"category"`
	Quantity        int32              `bson:"quantity"`
	CopiesAvailable int32              `bson:"capies_avaliable"`
}

func NewBook(name, author, category string, quantity, copies int32) *Book {
	return &Book{
		Name:            name,
		Author:          author,
		Category:        category,
		Quantity:        quantity,
		CopiesAvailable: copies,
	}
}
