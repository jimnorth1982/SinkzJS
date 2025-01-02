package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sinkzjs.org/m/v2/items/types"
)

type MongoDBProvider struct {
	Client       *mongo.Client
	DatabaseName string
	URI          string
}

func NewMongoDBProvider(databaseName string) *MongoDBProvider {
	return &MongoDBProvider{DatabaseName: databaseName}
}

func (p *MongoDBProvider) Connect(uri string, creds options.Credential, maxPoolSize uint64) error {
	if p.Client != nil {
		log.Printf("Already connected.")
		return nil
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(creds)
	clientOptions.SetMinPoolSize(maxPoolSize / 2)
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

func insertAllIntoColl(collection *mongo.Collection, items *[]interface{}) error {
	_, err := collection.InsertMany(context.TODO(), *items)
	return err
}

func (p *MongoDBProvider) GetItems() (*[]types.Item, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.Item](p, "items", filter, opts)
}

func (p *MongoDBProvider) GetItemById(id uint64) (*types.Item, error) {
	filter := bson.D{{Key: "id", Value: id}}
	sort := bson.D{}
	opts := options.FindOne().SetSort(sort)
	return FindOneFromCollection[types.Item](p, "items", filter, opts)
}

func (p *MongoDBProvider) AddItem(item *types.Item) (*types.Item, error) {
	coll, err := p.Collection("items")
	if err != nil {
		log.Fatalf("Error getting collection: [%s] Error: %v", "items", err)
		return nil, err
	}

	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}

	insertAllIntoColl(coll, &documents)
	return &types.Item{}, nil
}

func (p *MongoDBProvider) GetRarities() (*[]types.Rarity, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.Rarity](p, "rarity", filter, opts)
}

func (p *MongoDBProvider) GetItemTypes() (*[]types.ItemType, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.ItemType](p, "item_types", filter, opts)
}

func GetDefaultSortOptionsAndFilter() (bson.D, *options.FindOptions) {
	filter := bson.D{}
	sort := bson.D{}
	opts := options.Find().SetSort(sort)
	return filter, opts
}

func GetListFromCollection[T interface{}](provider *MongoDBProvider, collName string, filter bson.D, opts *options.FindOptions) (*[]T, error) {
	coll, err := provider.Collection(collName)
	if err != nil {
		log.Fatalf("Error finding collection: %s %v", collName, err)
		return nil, err
	}

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatalf("Error finding data in collection: %s %v", collName, err)
		return nil, err
	}
	var results []T
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func FindOneFromCollection[T interface{}](provider *MongoDBProvider, collName string, filter bson.D, opts *options.FindOneOptions) (*T, error) {
	collection, err := provider.Collection("items")
	if err != nil {
		log.Fatalf("Error finding collection: items %v", err)
		return nil, err
	}
	result := collection.FindOne(context.TODO(), filter, opts)
	if result == nil {
		log.Fatalf("Error finding item")
		return nil, err
	}

	var item T
	if err = result.Decode(&item); err != nil {
		log.Fatalf("Error decoding item: %v", err)
		return nil, err
	}

	return &item, nil
}

func (p *MongoDBProvider) GetImages() (*[]types.Image, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetAttributes() (*[]types.Attribute, error) {
	return nil, nil
}

func (p *MongoDBProvider) GetAttributeGroupings() (*[]types.AttributeGrouping, error) {
	return nil, nil
}

func (p *MongoDBProvider) ItemNameExistsInDb(string) bool {
	return true
}

func (p *MongoDBProvider) UpdateItem(uint64, *types.Item) (*types.Item, error) {
	return &types.Item{}, nil
}

func (p *MongoDBProvider) CleanupConnection() {
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := p.Client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

var _ Provider = (*MongoDBProvider)(nil)
