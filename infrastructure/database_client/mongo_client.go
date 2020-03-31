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
	Schema   = "MONGODB_SCHEMA"
)
var (
	Client *mongo.Client
	)

func init(){
	var err error
	mongoHost     := Host
	mongoPort, _  := strconv.ParseInt(Port, 10, 64)
	mongoUsername := os.Getenv(Username)
	mongoPassword := os.Getenv(Password)
	dataSource := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoUsername, mongoPassword, mongoHost, mongoPort)
	clientOptions := options.Client().ApplyURI(dataSource)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}
