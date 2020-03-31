package database_client

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client        *mongo.Client
	mongoHost     = os.Getenv("MONGODB_HOST")
	mongoPort, _  = strconv.ParseInt(os.Getenv("MONGODB_PORT"), 10, 64)
	mongoUsername = os.Getenv("MONGODB_USERNAME")
	mongoPassword = os.Getenv("MONGODB_PASSWORD")
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
