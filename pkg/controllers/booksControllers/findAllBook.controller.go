package booksControllers

import (
	"net/http"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/services/booksServices"
)

func FindAllBooks(rw http.ResponseWriter, r *http.Request) {
	books, err := booksServices.FindAllBooks()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, err2 := rw.Write([]byte("Err is in findALL"))
		if err2 != nil {
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(books)
}
