package booksServices

import (
	"encoding/json"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func FindAllBooks() ([]byte, error) {
	books, err := book.FindAllBooks()
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}
