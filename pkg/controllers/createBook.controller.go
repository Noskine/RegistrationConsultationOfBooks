package controllers

import (
	"io"
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services"
)

func CreateBook(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := services.CreateBookService(reqBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusCreated)
}
