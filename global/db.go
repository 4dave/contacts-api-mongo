package global

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectToMongo()
}

func connectToMongo() {
	DBURI := os.Getenv("DBURI")
	client, err := mongo.NewClient(options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	DB = *client.Database("contacts")
	fmt.Println("Connected to MongoDB")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// fmt.Println("connected to mongodb")
	// DB = *client.Database("contacts")
}
