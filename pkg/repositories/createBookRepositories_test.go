package repositories

import (
	"database/sql"
	"fmt"
	"testing"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
)

func TestCreateBookRepositories(t *testing.T) {
	var book *Book = NewBook("Querido Jhon", "Nicholas Spark", 3)
	var pass string = fmt.Sprintf("\033[33m\033[1m PASS...")
	var db *sql.DB = c.Connection()

	t.Run("Testando a criação de livros", func(t *testing.T) {
		t.Helper()
		esperado := &Book{
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

	t.Run("Testando a criação de tabelas no banco de dados", func(t *testing.T) {
		err := createTabaleIfNotExists(db)
		if err != nil {
			t.Fatalf("\u001B[31m Error na criação ou validação da tabela: \u001B[1m %v", err)
		}

		t.Log("\n" + "Testando a criação de tabelas no banco de dados" + pass)
	})

	t.Run("Testando a Inserção de Dados na tabela Books", func(t *testing.T) {
		err := insetInto(db, book)
		if err != nil {
			t.Fatalf("\u001B[31m Error na criação ou validação da tabela: \u001B[1m %v", err)
		}

		t.Log("\n" + "Testando a Inserção de Dados na tabela Books" + pass)
	})
}
