package storage

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sinkzjs.org/m/v2/items/types"
)

type MongoStorageProvider struct {
	Client       *mongo.Client
	DatabaseName string
	URI          string
	log          slog.Logger
}

type MongoQuery struct {
	provider *MongoStorageProvider
	collName string
	filter   bson.D
	opts     *options.FindOptions
}

func NewMongoStorageProvider(databaseName string) *MongoStorageProvider {
	return &MongoStorageProvider{
		DatabaseName: databaseName,
		log:          *slog.Default().With("area", "MongoDBProvider"),
	}
}

func (p *MongoStorageProvider) Connect(uri string, creds options.Credential, maxPoolSize uint64) error {
	if p.Client != nil {
		p.log.Warn("Already connected.")
		return nil
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(creds)
	clientOptions.SetMinPoolSize(maxPoolSize / 2)
	clientOptions.SetMaxPoolSize(maxPoolSize)
	client, err := connectToMongoDB(clientOptions)
	if err != nil {
		p.log.Error("Failed to connect to MongoDB")
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

func (p *MongoStorageProvider) GetItems() (*[]types.Item, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.Item](&MongoQuery{p, "items", filter, opts})
}

func (p *MongoStorageProvider) GetItemById(id uint64) (*types.Item, error) {
	filter := bson.D{{Key: "id", Value: id}}
	sort := bson.D{}
	opts := options.FindOne().SetSort(sort)
	return FindOneFromCollection[types.Item](p, "items", filter, opts)
}

func (p *MongoStorageProvider) AddItem(item *types.Item) (*types.Item, error) {
	coll, err := p.Collection("items")
	if err != nil {
		p.log.Error("Error getting collection: [%s] Error: %v", "items", err)
		return nil, err
	}

	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}

	insertAllIntoColl(coll, &documents)
	return &types.Item{}, nil
}

func (p *MongoStorageProvider) GetRarities() (*[]types.Rarity, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.Rarity](&MongoQuery{p, "rarity", filter, opts})
}

func (p *MongoStorageProvider) GetItemTypes() (*[]types.ItemType, error) {
	filter, opts := GetDefaultSortOptionsAndFilter()
	return GetListFromCollection[types.ItemType](&MongoQuery{p, "item_types", filter, opts})
}

func GetDefaultSortOptionsAndFilter() (bson.D, *options.FindOptions) {
	filter := bson.D{}
	sort := bson.D{}
	opts := options.Find().SetSort(sort)
	return filter, opts
}

func GetListFromCollection[T interface{}](query *MongoQuery) (*[]T, error) {
	coll, err := query.provider.Collection(query.collName)
	if err != nil {
		log.Fatalf("Error finding collection: %s %v", query.collName, err)
		return nil, err
	}

	cursor, err := coll.Find(context.TODO(), query.filter, query.opts)
	if err != nil {
		log.Fatalf("Error finding data in collection: %s %v", query.collName, err)
		return nil, err
	}
	var results []T
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func FindOneFromCollection[T interface{}](provider *MongoStorageProvider, collName string, filter bson.D, opts *options.FindOneOptions) (*T, error) {
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

func (p *MongoStorageProvider) GetImages() (*[]types.Image, error) {
	return nil, errors.New("Unsupported")
}

func (p *MongoStorageProvider) GetAttributes() (*[]types.Attribute, error) {
	return nil, errors.New("Unsupported")
}

func (p *MongoStorageProvider) GetAttributeGroupings() (*[]types.AttributeGrouping, error) {
	return nil, errors.New("Unsupported")
}

func (p *MongoStorageProvider) ItemNameExistsInDb(string) bool {
	return true
}

func (p *MongoStorageProvider) UpdateItem(uint64, *types.Item) (*types.Item, error) {
	return &types.Item{}, errors.New("Unsupported")
}

func (p *MongoStorageProvider) CleanupConnection() {
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := p.Client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

var _ StorageProvider = (*MongoStorageProvider)(nil)
