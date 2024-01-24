package dto

import "encoding/json"

type (
	InputRegisterUserDTO struct {
		UserName string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
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

func (o *OutputRegisterUserDTO) GetMessage() []byte {
	js, _ := json.Marshal(o.message)
	return js
}
