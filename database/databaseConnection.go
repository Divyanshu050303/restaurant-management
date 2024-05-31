package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	client, error := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if error != nil {
		log.Fatal(error)
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancle()

	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("rastaurant").Collection(collectionName)
	return collection

}
