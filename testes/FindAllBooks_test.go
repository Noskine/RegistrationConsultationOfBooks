package testes

import (
	"testing"

	"github.com/Noskine/RegistrationConsultationOfBooks/pkg/repositories/book"
)

func TestFindAllBooks(t *testing.T) {
	books, err := book.FindAllBooks()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*books)
}

func BenchmarkFindAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		books, err := book.FindAllBooks()
		if err != nil {
			b.Fatal(err)
		}

		b.Log(books)
	}
}
