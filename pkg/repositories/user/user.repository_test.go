package user

import (
	"testing"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
)

func TestUserRepository(t *testing.T) {
	db := c.Connection()

	t.Run("", func(t *testing.T) {
		err := ifNotExists(db)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("PASS")
	})
}
