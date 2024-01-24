package ihttp

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/controllers"
	"github.com/Noskine/RegistrationConsultationOfBooks/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	return &Server{
		e: echo.New(),
	}
}

func (s *Server) LoadServer() *echo.Echo {
	s.e.Validator = &utils.CustomValidator{
		V: validator.New(),
	}
	s.Routes()

	return s.e
}

func (s *Server) Routes() {
	g := s.e.Group("/api")

	g.POST("/public", controllers.RegisterUser)
}
