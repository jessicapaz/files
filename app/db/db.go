package db

import (
	"context"

	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientError error

var mongoOnce sync.Once

//Client returns a mongodb client
func Client() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientError = err
		}
		clientInstance = client
	})
	return clientInstance, clientError
}
