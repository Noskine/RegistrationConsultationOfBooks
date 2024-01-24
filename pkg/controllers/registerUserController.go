package controllers

import (
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/internal/usecases"
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/dto"
	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	input := new(dto.InputRegisterUserDTO)

	if err := c.Bind(input); err != nil {
		output := dto.NewOutputRegisterUserDTO("Bad Request")
		return c.JSONBlob(http.StatusBadRequest, output.GetMessage())
	}

	if err := c.Validate(input); err != nil {
		output := dto.NewOutputRegisterUserDTO("Validate Error!")
		return c.JSONBlob(http.StatusBadRequest, output.GetMessage())
	}

	message, err := usecases.RegisterUserUsecase(input)
	if err != nil {
		return c.JSONBlob(http.StatusBadRequest, message.GetMessage())
	}

	return c.JSONBlob(http.StatusOK, message.GetMessage())
}
