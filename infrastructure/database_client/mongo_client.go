package database_client

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Host     = "MONGODB_HOST"
	Port     = "MONGODB_PORT"
	//Username
	Username = "MONGODB_USERNAME"
	//Password
	Password = "MONGODB_PASSWORD"
)

func GetDatabaseInstance() *mongo.Client {
	var error error
	mongoHost     := os.Getenv(Host)
	mongoPort, _  := strconv.ParseInt(os.Getenv(Port), 10, 64)
	mongoUsername := os.Getenv(Username)
	mongoPassword := os.Getenv(Password)
	dataSource := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoUsername, mongoPassword, mongoHost, mongoPort)

	clientOptions := options.Client().ApplyURI(dataSource)
	var client  *mongo.Client
	client, error = mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		panic(error)
	}

	error = client.Ping(context.TODO(), nil)

	if error != nil {
		panic(error)
	}
	return client
}