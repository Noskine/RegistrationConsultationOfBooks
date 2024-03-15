package main

import (
	"net/http"

	ihttp "github.com/Noskine/RegistrationConsultationOfBooks/cmd/http"
	"github.com/Noskine/RegistrationConsultationOfBooks/config"
)

func main() {
	env := config.LoadEnvironments()
	e := ihttp.NewServer().LoadServer()

	if err := e.Start(env.APP_PORT); err != http.ErrServerClosed {

		e.Logger.Fatal(err)
	}
}
