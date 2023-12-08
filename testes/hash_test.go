package testes

import (
	"reflect"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/utils"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateHashPassword(t *testing.T) {
	password := []byte("1234523341435243126")

	t.Run("Testando se o retorno do password é igual a string", func(t *testing.T) {
		passwordHash := utils.PasswordHash(password)

		if reflect.TypeOf(passwordHash) != reflect.TypeOf("") {
			t.Fatal("o hash não retornou um valor do tipo string!")
		}

		t.Log("PASS")
	})

	t.Run("", func(t *testing.T) {
		passwordHash := utils.PasswordHash(password)
		err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
		if err != nil {
			t.Fatalf("os hash não são iguais: %v", err)
		}

		t.Log("PASS")
	})
}
