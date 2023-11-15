package testes

import (
	"fmt"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services"
)

func TestCreateBookService(t *testing.T) {
	var pass string = fmt.Sprintf("\033[33m\033[1m PASS...")

	t.Run("", func(t *testing.T) {
		err := services.CreateBookService([]byte(`
					{
						"name":"Um porto seguro",
						"author": "nicholas spark",
						"quantity": 2
					}
				`))
		if err != nil {
			t.Fatalf("Erro no serviço de criação dos livros")
		}

		t.Log("\n" + "Testando a Inserção de Dados na tabela Books" + pass)
	})
}
