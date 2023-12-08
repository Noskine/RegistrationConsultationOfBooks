package usersServices

import (
	"encoding/json"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/user"
)

func FindPyPK(id string) ([]byte, error) {
	us, err := user.FindByPk(id)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(us)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}
