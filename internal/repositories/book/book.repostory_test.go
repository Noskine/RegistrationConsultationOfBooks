package book_test

import (
	"encoding/json"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/repositories/book"
)

func TestBookRepository(t *testing.T) {
	b := entities.NewBook("Querido John", "Nicholas Spark", 1)
	t.Run("Testing the creation of the books", func(t *testing.T) {
		err := book.Create(b)
		if err != nil {
			t.Errorf("Error creating the book -> %s", err)
			return
		}

		t.Log("PASS! Book Creation Test")
	})

	t.Run("Testing to search for all books in the database", func(t *testing.T) {
		res, err := book.FindAll()
		if err != nil {
			t.Errorf("Error trying to fetch all books -> %s", err)
		}

		sBytes, err := json.Marshal(res)
		if err != nil {
			t.Errorf("Error when transforming the struct book into a byte slice-> %s", err)
		}

		t.Logf("Searches for the books completed successfully! -> %s", string(sBytes))
	})

	t.Run("", func(t *testing.T) {
		id := "3c872316-40fd-4995-b67e-aae6a8df3ac6"

		res, err := book.FindByPk(id)
		if err != nil {
			t.Errorf("Error searching for book with primary key -> %s", err)
		}

		sBytes, err := json.Marshal(res)
		if err != nil {
			t.Errorf("Error when transforming the struct book into a byte slice-> %s", err)
		}

		t.Logf("Success in searching for the book via primary key -> %s", string(sBytes))
	})
}
