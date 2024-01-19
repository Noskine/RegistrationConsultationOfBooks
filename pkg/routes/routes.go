package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	g := e.Group("/api")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá mundo")
	})

	g.GET("/o", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá mundo")
	})
}
