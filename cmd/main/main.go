package main

import (
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/config"
	ihttp "github.com/Noskine/RegistrationConsultationOfBooks/pkg/http"
)

func main() {
	env := config.LoadEnvironments()
	e := ihttp.NewServer().LoadServer()

	if err := e.Start(env.APP_PORT); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
