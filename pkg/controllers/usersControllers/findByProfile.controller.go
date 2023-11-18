package usersControllers

import (
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services/usersServices"
	"github.com/gorilla/mux"
)

func FindPyPkUsers(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err2 := rw.Write([]byte("Err is in findByPk"))
		if err2 != nil {
			return
		}
	}

	profile, err := usersServices.FindPyPK(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, err2 := rw.Write([]byte("Err is in findByPK"))
		if err2 != nil {
			return
		}
	}

	_, _ = rw.Write(profile)
}
