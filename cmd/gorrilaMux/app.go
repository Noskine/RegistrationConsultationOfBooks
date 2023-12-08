package gorrilaMux

import (
	"log"
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/routes"
	"github.com/gorilla/mux"
)

func App() {
	app := mux.NewRouter()

	routes.Routes(app)

	log.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":3031", app))
}
