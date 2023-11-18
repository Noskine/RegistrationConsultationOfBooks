package testes

import (
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func TestFindAllBooks(t *testing.T) {
	t.Run("test de findall", func(t *testing.T) {
		books, err := book.FindAllBooks()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(*books)
	})
	t.Run("test de findbyPK", func(t *testing.T) {
		b, err := book.FindByPk("752317fb-b814-456d-8bfd-46ea929cc71e")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(*b)
	})
}
