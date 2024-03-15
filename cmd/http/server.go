package ihttp

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echo: echo.New(),
	}
}

func (s *Server) LoadServer() *echo.Echo {
	r := s.echo
	s.Routes()

	r.Validator = &utils.CustomValidator{
		V: validator.New(),
	}
	return r
}
