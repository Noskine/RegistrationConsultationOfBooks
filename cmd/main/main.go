package main

import (
	"log"
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	routes.Routes(app)

	log.Println("Server is running on port 3031")
	log.Fatal(http.ListenAndServe(":3031", app))
}
