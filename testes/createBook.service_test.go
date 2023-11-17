package testes

import (
	"fmt"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services"
)

func TestCreateBookService(t *testing.T) {
	var pass string = fmt.Sprintf("\033[33m\033[1m PASS...")

	err := services.CreateBookService([]byte(`
					{
						"name":"A linguagem de progamação go",
						"author": "",
						"quantity": 10
					}
				`))
	if err != nil {
		t.Fatalf("Erro no serviço de criação dos livros")
	}

	t.Log("\n" + "Testando a Inserção de Dados na tabela Books" + pass)
}
