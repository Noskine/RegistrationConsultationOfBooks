package user

import (
	"database/sql"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/Noskine/RegistrationConsultationOfBooks/utils"
	"github.com/google/uuid"
)

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_Hash string `json:"-"`
	function      string `json:"function"`
}

func NewUser(username, email, function, password_hash string) *User {
	return &User{
		Id:            uuid.New().String(),
		Username:      username,
		Email:         email,
		Password_Hash: password_hash,
		function:      function,
	}
}

func ifNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
    	id varchar(80) NOT NULL, 
		username varchar(80) NOT NULL,
		password varchar(80), 
    	email varchar(80) NOT NULL UNIQUE,
		function varchar(80), 
		primary key (id)
    );`)
	if err != nil {
		return err
	}
	return nil
}

func insetInto(db *sql.DB, user *User) error {
	inset, err := db.Prepare("INSERT INTO users(id, username, password, email, function) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = inset.Exec(user.Id, user.Username, user.Password_Hash, user.Email, user.function)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(username, email, password, function string) error {
	db := c.Connection()
	defer db.Close()

	if err := ifNotExists(db); err != nil {
		return err
	}

	password_hash := utils.PasswordHash([]byte(password))

	user := NewUser(username, email, function, password_hash)

	err := insetInto(db, user)
	if err != nil {
		return err
	}

	return nil
}
