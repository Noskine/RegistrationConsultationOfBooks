package book

import (
	"database/sql"
	"log"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/google/uuid"
)

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int8   `json:"quantity"`
}

func NewBook(name, author string, quantity int8) *Book {
	return &Book{
		Id:       uuid.New().String(),
		Name:     name,
		Author:   author,
		Quantity: quantity,
	}
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

func CreateBook(name, author string, quantity int8) error {
	var db *sql.DB = c.Connection()
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
