package db

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
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

func (p *MongoDBProvider) Database() (*mongo.Database, error) {
	if p.Client == nil {
		return nil, mongo.ErrClientDisconnected
	}

	return p.Client.Database(p.DatabaseName), nil
}

func (p *MongoDBProvider) Collection() (*mongo.Collection, error) {
	if p.Client == nil {
		return nil, mongo.ErrClientDisconnected
	}

	return p.Client.Database(p.DatabaseName).Collection(p.CollectionName), nil
}

func (p *MongoDBProvider) ClearAndLoadDataFromJSON() error {
	// Load items from JSON file
	items, err := loadItemsFromFile("/home/jimi/dev/SinkzJS/backend/items/db/data/item_data.json")
	if err != nil {
		log.Fatalf("Failed to load items from file: %v", err)
		return err
	}

	collection := p.Client.Database("items").Collection("items")
	// Drop the collection
	err = collection.Drop(context.Background())
	if err != nil {
		log.Fatalf("Failed to drop: %v", err)
		return err
	}

	log.Println("Collection cleared successfully.")

	err = insertItemsIntoMongoDB(collection, &items)
	if err != nil {
		log.Fatalf("Failed to insert items into MongoDB: %v", err)
		return err
	}
	log.Println("Items successfully inserted into MongoDB")
	return nil
}
