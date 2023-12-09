package user

import (
	"github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
)

func createTableAndIfNotExist() error {
	db := config.Connection()

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users 
		(id varchar(80) NOT NULL, 
		email varchar(80) NOT NULL UNIQUE, 
		password_hash varchar(255) NOT NULL, 
		username varchar(80) NOT NULL,
		fn varchar(80), 
		primary key (id)
	);
`)
	if err != nil {
		return err
	}
	return nil
}

func Create(u *entities.User) error {
	err := createTableAndIfNotExist()
	if err != nil {
		return err
	}
	db := config.Connection()

	stmt, err := db.Prepare("INSERT INTO `users` (`id`, `email`,`password_hash`,`username`,`fn`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = stmt.Exec(u.Id, u.Email, u.Password, u.Username, u.Function)
	if err != nil {
		return err
	}

	return nil
}

func FindByPk(id string) (*entities.User, error){
	db := config.Connection()

	stmt, err := db.Prepare("SELECT * FROM users WHERE id =?")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var u entities.User
	stmt.QueryRow(id).Scan(u.Id, u.Email, u.Password, u.Username, u.Function)

	return &u, nil
}