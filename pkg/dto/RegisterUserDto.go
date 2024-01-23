package dto

type (
	InputRegisterUserDTO struct {
		UserName string
		Email    string
	}

	OutputRegisterUserDTO struct {
		message string
	}
)

func NewInputRegisterUserDTO(user, email string) *InputRegisterUserDTO {
	return &InputRegisterUserDTO{
		UserName: user,
		Email:    email,
	}
}

func NewOutputRegisterUserDTO(message string) *OutputRegisterUserDTO {
	return &OutputRegisterUserDTO{message}
}
