package usersControllers

import (
	"io"
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services/usersServices"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := usersServices.CreateUserService(reqBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusCreated)
}
