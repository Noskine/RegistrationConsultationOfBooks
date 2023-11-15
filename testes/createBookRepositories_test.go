package testes

import (
	"fmt"
	"testing"

	book2 "github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func TestCreateBookRepositories(t *testing.T) {
	var book *book2.Book = book2.NewBook("Querido Jhon", "Nicholas Spark", 3)
	var pass string = fmt.Sprintf("\033[33m\033[1m PASS...")

	t.Run("Testando a criação de livros", func(t *testing.T) {
		t.Helper()
		esperado := &book2.Book{
			Id:       book.Id,
			Name:     "Querido Jhon",
			Author:   "Nicholas Spark",
			Quantity: 3,
		}

		if *book != *esperado {
			t.Errorf("\u001B[31m o resultado não foi como o esperado! [%v]", esperado)
			return
		}

		t.Log(pass)
	})
}
