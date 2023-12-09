package book

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
)

func createTableAndIfNotExist() error {
	db := config.Connection()

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS books
		(id varchar(80) NOT NULL, 
		name varchar(80) NOT NULL, 
		author varchar(80), 
		quantity integer DEFAULT 1, 
		primary key (id)
	);
`)
	if err != nil {
		return err
	}
	return nil
}

func Create(b *entities.Book) error {
	db := config.Connection()

	err := createTableAndIfNotExist()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO `books` (`id`, `name`,`author`,`quantity`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = stmt.Exec(b.Id, b.Name, b.Author, b.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func FindAll() (*[]entities.Book, error) {
	db := config.Connection()

	rows, err := db.Query("SELECT * FROM books ORDER BY name;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var b []entities.Book

	for rows.Next() {
		var book entities.Book
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

func FindByPk(id string) (*entities.Book, error) {
	db := config.Connection()

	stmt, err := db.Prepare("SELECT * FROM books WHERE id =?")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var b entities.Book
	stmt.QueryRow(id).Scan(&b.Id, &b.Name, &b.Author, &b.Quantity)

	return &b, nil
}
