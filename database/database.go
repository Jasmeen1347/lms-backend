package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	url := os.Getenv("MONGO_URL")
	fmt.Println(url)
	client, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	DB = client.Database("lmswithgo")
	fmt.Println("Connected to database")
}
