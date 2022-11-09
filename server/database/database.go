package database

import (
	"context"
	"fmt"
	"log"

	"github.com/johnnyaustor/golang-skeleton-nosql/server/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Connection *mongo.Client
	Name       string
}

func ConnectDatabase() *DB {
	log.Println("Connect to database " + infoDB())

	uri := configuration.Conf.Database.Uri
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(fmt.Sprintf("Cannot Connect Database %v", err))
	}

	return &DB{Connection: client, Name: "welcome"}
}

func infoDB() string {
	return fmt.Sprintf("%v", configuration.Conf.Database.Uri)
}
