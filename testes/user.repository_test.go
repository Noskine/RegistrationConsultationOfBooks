package testes

import (
	"testing"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/user"
)

func TestUserRepository(t *testing.T) {
	db := c.Connection()

	t.Run("", func(t *testing.T) {
		err := user.ifNotExists(db)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("PASS")
	})
}
