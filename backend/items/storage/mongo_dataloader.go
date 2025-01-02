package storage

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"sinkzjs.org/m/v2/items/types"
)

func loadFromFile(filename string) ([]types.Item, error) {
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

func (p *MongoStorageProvider) Database() (*mongo.Database, error) {
	if p.Client == nil {
		return nil, mongo.ErrClientDisconnected
	}

	return p.Client.Database(p.DatabaseName), nil
}

func (p *MongoStorageProvider) Collection(collName string) (*mongo.Collection, error) {
	if p.Client == nil {
		return nil, mongo.ErrClientDisconnected
	}

	return p.Client.Database(p.DatabaseName).Collection(collName), nil
}

func (p *MongoStorageProvider) ClearAndLoadDataFromJSON() error {
	items, err := loadFromFile("/home/jimi/dev/SinkzJS/backend/items/storage/data/item_data.json")
	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	itemsList := make([]interface{}, 0, len(items))
	var rarities = map[string]types.Rarity{}
	for _, item := range items {
		itemsList = append(itemsList, item)
		rarities[item.Rarity.Name] = item.Rarity
	}

	rarityList := make([]interface{}, 0, len(rarities))
	for _, rarity := range rarities {
		rarityList = append(rarityList, rarity)
	}

	if err := dropCollection(p, "items"); err != nil {
		p.log.Error(err.Error())
		return err
	}

	if err := AddElementsToCollection(p, "items", itemsList); err != nil {
		p.log.Error(err.Error())
		return err
	}

	if err := dropCollection(p, "rarity"); err != nil {
		p.log.Error(err.Error())
		return err
	}

	if err := AddElementsToCollection(p, "rarity", rarityList); err != nil {
		p.log.Error(err.Error())
		return err
	}

	return nil
}

func dropCollection(p *MongoStorageProvider, collName string) error {
	if coll, err := p.Collection(collName); err != nil {
		p.log.Error(err.Error())
		return err
	} else {
		p.log.Info("Dropping collection " + collName)
		if err := coll.Drop(context.Background()); err != nil {
			p.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func AddElementsToCollection(p *MongoStorageProvider, collName string, elements []interface{}) error {
	coll, err := p.Collection(collName)
	if err != nil {
		p.log.Error(err.Error())
		return err
	}

	p.log.Info("Adding " + strconv.Itoa(len(elements)) + " documents to collection " + collName)
	_, err = coll.InsertMany(context.TODO(), elements)
	return err
}
