package db

import (
	"context"
	"time"

	"github.com/Noskine/RegistrationConsultationOfBooks/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	client *mongo.Client
}

func NewConnection() *Db {
	env := config.LoadEnvironments()

	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(c, options.Client().ApplyURI(env.URI_DB))
	if err != nil {
		panic(err)
	}

	return &Db{
		client: client,
	}
}

func (db *Db) Database(name string) *mongo.Database {
	return db.client.Database(name)
}

func (db *Db) Desc() {
	db.client.Disconnect(context.Background())
}
