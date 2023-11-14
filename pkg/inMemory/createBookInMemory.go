package inMemory

import (
	"github.com/google/uuid"
)

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var itens []Book

func newBook(name, author string) *Book {
	return &Book{
		uuid.New().String(),
		name,
		author,
	}
}

func insetBook(b *Book) []Book {
	itens = append(itens, Book{
		b.Id,
		b.Name,
		b.Author,
	})

	return itens
}

func CreateBookInMemory(name, author string) []Book {
	book := newBook(name, author)
	return insetBook(book)
}
