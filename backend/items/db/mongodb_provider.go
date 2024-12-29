package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBProvider struct {
}

func NewMongoDBProvider() *MongoDBProvider {
	return &MongoDBProvider{}
}

func (p *MongoDBProvider) init() {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("No MONGODB_URI. Check environment.")
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (p *MongoDBProvider) GetConnection(uri string, userName string, password string) (*mongo.Client, error) {
	log.Printf("Logging in with URI: %s", uri)
	credentials := options.Credential{Username: userName, Password: password}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)

	// Connect to MongoDB
	client, err := connectToMongoDB(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	return client, nil
}
