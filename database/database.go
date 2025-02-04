package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance holds the MongoDB client and collection reference.
type MongoInstance struct {
	Client          *mongo.Client
	MongoCollection *mongo.Collection
}

// MI is a global instance for MongoDB connection.
var MI MongoInstance

// ConnectMongo connects to the MongoDB instance.
func ConnectMongo() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating Mongo client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	// Choose the database and collection
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "urlshortener"
	}
	collectionName := os.Getenv("MONGO_COLLECTION")
	if collectionName == "" {
		collectionName = "urls"
	}
	MI = MongoInstance{
		Client:          client,
		MongoCollection: client.Database(dbName).Collection(collectionName),
	}
	log.Println("Connected to MongoDB!")
}
