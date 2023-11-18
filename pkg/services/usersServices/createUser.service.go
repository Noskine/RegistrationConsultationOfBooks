package usersServices

import (
	"encoding/json"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/user"
)

func CreateUserService(sliceBytes []byte) error {
	primaryUser := user.User{}
	err := json.Unmarshal(sliceBytes, &primaryUser)
	if err != nil {
		return err
	}

	if err := user.CreateUser(primaryUser.Username, primaryUser.Email, primaryUser.Password, primaryUser.Function); err != nil {
		return err
	}

	return nil
}
