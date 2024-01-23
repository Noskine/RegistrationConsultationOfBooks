package ihttp

import "github.com/labstack/echo/v4"

type Router struct {
	E *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		E: echo.New(),
	}
}
