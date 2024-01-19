package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() (*mongo.Client, error) {
	env := LoadEnvironments()
	
	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(c, options.Client().ApplyURI(env.URI_DB))
	if err != nil {
		return nil, err
	}

	client.Database(LoadEnvironments().DATABASE)

	return client, nil
}
