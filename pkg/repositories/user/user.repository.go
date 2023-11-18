package user

import (
	"database/sql"

	c "github.com/Noskine/RegistrationConsultationOfBooks/config"
	"github.com/Noskine/RegistrationConsultationOfBooks/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Function string `json:"function"`
}

func NewUser(username, email, function, password_hash string) *User {
	return &User{
		Id:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: password_hash,
		Function: function,
	}
}

func ifNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS usersControllers (
    	id varchar(80) NOT NULL, 
		username varchar(80) NOT NULL,
		password varchar(80), 
    	email varchar(80) NOT NULL UNIQUE,
		position varchar(80) NOT NULL, 
		primary key (id)
    );`)
	if err != nil {
		return err
	}
	return nil
}

func insetInto(db *sql.DB, user *User) error {
	inset, err := db.Prepare("INSERT INTO usersControllers(id, username, password, email, position) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = inset.Exec(user.Id, user.Username, user.Password, user.Email, user.Function)
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

	passwordHash := utils.PasswordHash([]byte(password))

	user := NewUser(username, email, function, passwordHash)

	err := insetInto(db, user)
	if err != nil {
		return err
	}

	return nil
}

func FindByPk(id string) (*User, error) {
	db := c.Connection()
	stmt, err := db.Prepare("SELECT id, email, username, position FROM usersControllers WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	stmt.QueryRow(id).Scan(&user.Id, &user.Email, &user.Username, &user.Function)

	return &user, nil
}
