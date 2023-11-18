package routes

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/controllers/booksControllers"
	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/controllers/usersControllers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/createbook", booksControllers.CreateBook).Methods("POST")
	r.HandleFunc("/createuser", usersControllers.CreateUser).Methods("POST")
	r.HandleFunc("/book", booksControllers.FindAllBooks).Methods("GET")
	r.HandleFunc("/bookbypk/{id}", booksControllers.FindPyPKBook).Methods("GET")
	r.HandleFunc("/profile/{id}", usersControllers.FindPyPkUsers).Methods("GET")
}
