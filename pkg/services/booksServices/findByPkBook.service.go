package booksServices

import (
	"encoding/json"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func FindPyPKBook(id string) ([]byte, error) {
	b, err := book.FindByPk(id)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}
