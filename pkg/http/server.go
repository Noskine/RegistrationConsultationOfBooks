package ihttp

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/controllers"
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
	return s.e
}

func (s *Server) Routes() {
	g := s.e.Group("/api")

	g.GET("/public", controllers.RegisterUser)
}
