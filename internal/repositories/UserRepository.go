package repositories

import (
	"context"

	db "github.com/Noskine/RegistrationConsultationOfBooks/config/database"
	"github.com/Noskine/RegistrationConsultationOfBooks/internal/entities"
)

func RegisterUser(user *entities.User) error {
	client := db.NewConnection()

	defer client.Desc()

	coll := client.Database("School").Collection("User")

	_, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}
