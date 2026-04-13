package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() *mongo.Database {
	uri := os.Getenv("MONGO_LOCAL_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	if uri == "" || dbName == "" {
		log.Fatal("Missing MongoDB environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ MongoDB Connected")

	DB = client.Database(dbName)

	initIndexes()

	return DB
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}

func initIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := DB.Collection("urls")

	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]int{"short_code": 1},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		log.Println("Index error:", err)
	}
}
