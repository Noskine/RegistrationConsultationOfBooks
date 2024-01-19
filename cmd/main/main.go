package main

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/config"
	ihttp "github.com/Noskine/RegistrationConsultationOfBooks/pkg/http"
	r "github.com/Noskine/RegistrationConsultationOfBooks/pkg/routes"
)

func main() {
	env := config.LoadEnvironments()

	e := ihttp.NewRouter().E
	r.Routes(e)
	
	e.Logger.Fatal(e.Start(env.APP_PORT))
}
