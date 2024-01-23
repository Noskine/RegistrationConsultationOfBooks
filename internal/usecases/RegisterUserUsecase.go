package usecases

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/repositories"
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/dto"
)

func RegisterUserUsecase(req *dto.InputRegisterUserDTO) (*dto.OutputRegisterUserDTO, error) {
	user := entities.NewUser(req.UserName, req.Email)

	if err := repositories.RegisterUser(user); err != nil {
		return nil, err
	}

	return dto.NewOutputRegisterUserDTO("Error ao cadrastrar o usuário"), nil
}
