package repositories

import (
	"database/sql"

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

func createTabaleIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books 
		(id varchar(80), 
		name varchar(80), 
		author varchar(80), 
		quantity integer, 
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

func CreateBookRepositories(name, author string, quantity int8) error {
	db := c.Connection()
	println(db)
	err := createTabaleIfNotExists(db)
	if err != nil {
		return err
	}
	book := NewBook(name, author, quantity)

	err = insetInto(db, book)
	if err != nil {
		return err
	}

	return nil
}
