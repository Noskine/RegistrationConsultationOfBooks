package ihttp

import "github.com/labstack/echo/v4"

func (server *Server) Routes() {
	g := server.echo.Group("/api")

	g.GET("/", func(c echo.Context) error {
		return c.JSONBlob(200, []byte("Olá mundo!"))
	})
}
