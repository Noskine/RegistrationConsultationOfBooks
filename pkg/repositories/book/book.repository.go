package book

import (
	"database/sql"
	"log"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/google/uuid"
)

var db *sql.DB = c.Connection()

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int8   `json:"quantity"`
}

type Books []Book

func NewBook(name, author string, quantity int8) *Book {
	return &Book{
		Id:       uuid.New().String(),
		Name:     name,
		Author:   author,
		Quantity: quantity,
	}
}

func CreateBook(name, author string, quantity int8) error {
	defer db.Close()
	if err := ifNotExists(db); err != nil {
		return err
	}
	book := NewBook(name, author, quantity)

	if err := insetInto(db, book); err != nil {
		return err
	}

	log.Println("Inserção no banco de dados conluida com sucesso!")
	return nil
}

func FindAllBooks() (*Books, error) {
	defer db.Close()
	rows, err := db.Query("SELECT * FROM books ORDER BY name;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var b Books
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Quantity)
		if err != nil {
			return nil, err
		}
		b = append(b, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func ifNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books 
		(id varchar(80) NOT NULL, 
		name varchar(80) NOT NULL, 
		author varchar(80), 
		quantity integer DEFAULT 1, 
		primary key (id));
`)
	if err != nil {
		return err
	}
	return nil
}

func insetInto(db *sql.DB, b *Book) error {
	stmt, err := db.Prepare("INSERT INTO books(id, name, author, quantity) VALUES( ?, ?, ?, ? ) ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Id, b.Name, b.Author, b.Quantity)
	if err != nil {
		return err
	}

	return nil
}
