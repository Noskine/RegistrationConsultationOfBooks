package controllers

import (
	"io"
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := services.CreateUserService(reqBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusCreated)
}
