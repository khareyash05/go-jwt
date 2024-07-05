package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbinstance() *mongo.Client {
	mongoURI := "mongodb://jwtMongo:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

// Client Database instance
var Client *mongo.Client = Dbinstance()

func OpenCollection(client *mongo.Client, collecionName string) *mongo.Collection {
	databaseName := "myDatabase"
	var collection *mongo.Collection = client.Database(databaseName).Collection(collecionName)
	return collection

}
