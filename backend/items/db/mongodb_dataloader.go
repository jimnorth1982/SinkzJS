package db

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sinkzjs.org/m/v2/items/types"
)

func loadItemsFromFile(filename string) ([]types.Item, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var items []types.Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func connectToMongoDB(clientOptions *options.ClientOptions) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func insertItemsIntoMongoDB(collection *mongo.Collection, items []types.Item) error {
	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}

	_, err := collection.InsertMany(context.TODO(), documents)
	return err
}

func ClearAndLoadDataFromSJON(client *mongo.Client) error {
	// Load items from JSON file
	items, err := loadItemsFromFile("/home/jimi/dev/SinkzJS/backend/items/db/data/item_data.json")
	if err != nil {
		log.Fatalf("Failed to load items from file: %v", err)
		return err
	}

	collection := client.Database("items").Collection("items")
	// Drop the collection
	err = collection.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Collection cleared successfully.")
	// Insert items into MongoDB
	err = insertItemsIntoMongoDB(collection, items)
	if err != nil {
		log.Fatalf("Failed to insert items into MongoDB: %v", err)
		return err
	}

	log.Println("Items successfully inserted into MongoDB")
	return nil
}
