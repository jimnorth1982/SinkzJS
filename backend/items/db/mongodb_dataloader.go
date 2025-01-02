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

func (p *MongoDBProvider) Collection(collName string) (*mongo.Collection, error) {
	if p.Client == nil {
		return nil, mongo.ErrClientDisconnected
	}

	return p.Client.Database(p.DatabaseName).Collection(collName), nil
}

func (p *MongoDBProvider) ClearAndLoadDataFromJSON() error {
	// Load items from JSON file
	items, err := loadItemsFromFile("/home/jimi/dev/SinkzJS/backend/items/db/data/item_data.json")
	if err != nil {
		log.Fatalf("Failed to load items from file: %v", err)
		return err
	}

	itemsList := make([]interface{}, 0, len(items))
	var rarities = map[string]types.Rarity{}
	for _, item := range items {
		itemsList = append(itemsList, item)
	}

	if err := AddAllItemsToCollection(p, "items", itemsList); err != nil {
		log.Fatalf("cannot add items to database: %v", err)
		return err
	}

	for _, item := range items {
		rarities[item.Rarity.Name] = item.Rarity
	}

	rarityList := make([]interface{}, 0, len(rarities))
	for _, rarity := range rarities {
		rarityList = append(rarityList, rarity)
	}

	if err := AddAllItemsToCollection(p, "rarity", rarityList); err != nil {
		log.Fatalf("cannot add rarities to database: %v", err)
		return err
	}

	return nil
}

func AddAllItemsToCollection(p *MongoDBProvider, collName string, elements []interface{}) error {
	coll, err := p.Collection(collName)
	if err != nil {
		log.Fatalf("Failed to get collection rarity %v", err)
		return err
	}

	log.Printf("Dropping collection %s.", collName)
	err = coll.Drop(context.Background())
	if err != nil {
		log.Fatalf("Failed to drop %s: %v", collName, err)
		return err
	}

	log.Printf("Adding %d documents to collection %s.", len(elements), collName)
	_, err = coll.InsertMany(context.TODO(), elements)
	return err
}
