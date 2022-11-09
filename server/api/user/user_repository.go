package user

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/johnnyaustor/golang-skeleton-nosql/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Get() string
	InsertOne() string
}

type repository struct {
	*database.DB
	collection *mongo.Collection
}

func NewRepository(db *database.DB) Repository {
	return &repository{db, db.Connection.Database(db.Name).Collection("users")}
}

func (r *repository) Get() string {
	cur := r.collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "Agus"}})
	var result bson.M
	return fmt.Sprintf("%v", cur.Decode(&result))
}

func (r *repository) InsertOne() string {
	doc := bson.D{{Key: "name", Value: rand.Int()}}
	result, err := r.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return "error " + err.Error()
	}
	return fmt.Sprintf("Inserted Document with _id: %v", result.InsertedID)
}
