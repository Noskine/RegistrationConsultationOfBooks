package routes

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/controllers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/createbook", controllers.CreateBook).Methods("Post")
}
