package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sinkzjs.org/m/v2/items/types"
)

type MongoDBProvider struct {
	Client         *mongo.Client
	CollectionName string
	DatabaseName   string
	URI            string
}

func NewMongoDBProvider(databaseName string, collectionName string) *MongoDBProvider {
	return &MongoDBProvider{DatabaseName: databaseName, CollectionName: collectionName}
}

func (p *MongoDBProvider) Connect(uri string, creds options.Credential, maxPoolSize uint64) error {
	if p.Client != nil {
		log.Printf("Already connected.")
		return nil
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(creds)
	clientOptions.SetMaxPoolSize(maxPoolSize)
	client, err := connectToMongoDB(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return err
	}
	p.Client = client

	return nil
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

func insertItemsIntoMongoDB(collection *mongo.Collection, items *[]types.Item) error {
	var documents []interface{}
	for _, item := range *items {
		documents = append(documents, item)
	}

	_, err := collection.InsertMany(context.TODO(), documents)
	return err
}

func (p *MongoDBProvider) GetItems() (*[]types.Item, error) {
	filter := bson.D{}
	sort := bson.D{}
	opts := options.Find().SetSort(sort)
	collection, err := p.Collection()
	if err != nil {
		log.Fatalf("Error finding collection: %s", p.CollectionName)
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatalf("Error finding collection: %s", p.CollectionName)
		return nil, err
	}
	var results []types.Item
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (p *MongoDBProvider) GetItemById(uint64) (*types.Item, error) {
	return &types.Item{}, nil
}

func (p *MongoDBProvider) AddItem(*types.Item) (*types.Item, error) {
	return &types.Item{}, nil
}

func (p *MongoDBProvider) GetRarities() (*map[uint64]types.Rarity, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetItemTypes() (*map[uint64]types.ItemType, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetImages() (*map[uint64]types.Image, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetAttributes() (*map[uint64]types.Attribute, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetAttributeGroupings() (*map[uint64]types.AttributeGrouping, error) {
	return nil, nil
}

func (p *MongoDBProvider) ItemNameExistsInDb(string) bool {
	return true
}

func (p *MongoDBProvider) UpdateItem(uint64, *types.Item) (*types.Item, error) {
	return &types.Item{}, nil
}

var _ Provider = (*MongoDBProvider)(nil)
