package services

import (
	"encoding/json"
	"fmt"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func CreateBookService(sliceBytes []byte) error {
	primaryBook := book.Book{}

	if err := json.Unmarshal(sliceBytes, &primaryBook); err != nil {
		return err
	}

	if err := book.CreateBook(primaryBook.Name, primaryBook.Author, primaryBook.Quantity); err != nil {
		return fmt.Errorf("erro no cadastramento dos livros: %v", err)
	}

	return nil

}
