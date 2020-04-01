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
var (
	Client        *mongo.Client
	mongoHost     = os.Getenv(Host)
	mongoPort, _  = strconv.ParseInt(os.Getenv(Port), 10, 64)
	mongoUsername = os.Getenv(Username)
	mongoPassword = os.Getenv(Password)
)

func init() {
	var error error
	dataSource := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoUsername, mongoPassword, mongoHost, mongoPort)
	clientOptions := options.Client().ApplyURI(dataSource)

	Client, error = mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		panic(error)
	}

	error = Client.Ping(context.TODO(), nil)

	if error != nil {
		panic(error)
	}
}