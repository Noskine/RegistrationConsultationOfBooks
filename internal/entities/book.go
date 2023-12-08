package entities

import "github.com/google/uuid"

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func NewBook(name, author string, qnt int) *Book {
	return &Book{
		Id:       uuid.New().String(),
		Name:     name,
		Author:   author,
		Quantity: qnt,
	}
}
