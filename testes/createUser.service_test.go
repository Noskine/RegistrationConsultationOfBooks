package testes

import (
	"fmt"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services"
)

func TestName(t *testing.T) {
	// Testando a criação dos usuários

	var pass string = fmt.Sprintf("\033[33mPASS...")

	t.Run("Testando a criação de um usuário no banco de dados", func(t *testing.T) {
		err := services.CreateUserService([]byte(`{
			"username":"adm",
			"email": "adm@email.com",
			"password": "12345",
			"function": "GodUser" 
		}`))

		if err != nil {
			t.Fatal("\u001B[31m" + err.Error())
		}

		t.Log("\n" + pass + "Testando a Inserção de Dados na tabela 'Users'")
	})

	t.Run("", func(t *testing.T) {
		err := services.CreateUserService([]byte(`{
			"username":"adm",
			"email": "adm@email.com",
			"password": "12345",
			"function": "GodUser" 
		}`))

		if err == nil {
			t.Fatal("\u001B[31m" + "Não foi o resultado esperado...")
		}

		t.Log("\n" + pass + "Não possível passar dois parameters iguais na row 'Email'")
	})
}
