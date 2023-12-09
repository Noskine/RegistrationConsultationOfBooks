package user_test

import (
	"encoding/json"
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/repositories/user"
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/utils"
)

func TestUserRepository(t *testing.T) {
	var password_hash string = utils.PasswordHash([]byte("jao312"))
	u := entities.NewUser("João Pitomba", "jao_p011@outldook.com", "admin", password_hash)

	t.Run("Testing create user", func(t *testing.T) {
		err := user.Create(u)
		if err != nil {
			t.Error(err)
		}

		t.Log("Pass! user created!")
	})

	t.Run("Testing get user profile", func(t *testing.T) {
		res, err := user.FindByPk("")
		if err != nil {
			t.Error(err)
		}

		jso, err := json.Marshal(res)
		if err != nil {
			t.Errorf("Error transforming data into json. %s", err)
		}

		t.Logf("PASS! -> %s", string(jso))
	})

}
