package storage

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"sinkzjs.org/m/v2/items/types"
)

type SqLiteStorageProvider struct {
	DBFileName    string
	TimeoutMillis int
	TTLMillis     int
	sqlitedb      *sql.DB
}

func NewSqlite3Provider(DBFileName string, TimeoutMillis int, TTLMillis int) *SqLiteStorageProvider {
	conn := Connect(DBFileName)
	provider := &SqLiteStorageProvider{DBFileName, TimeoutMillis, TTLMillis, conn}
	provider.init()
	return provider
}

func (p *SqLiteStorageProvider) init() {
	log.Print("Init: SqLiteProvider")
}

func Connect(dbpath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatalf("DB Handle Error: %s", err.Error())
	}
	db.SetMaxOpenConns(5)
	return db
}

func (p *SqLiteStorageProvider) GetItems() (*[]types.Item, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetItemById(uint64) (*types.Item, error) {
	return &types.Item{}, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) AddItem(*types.Item) (*types.Item, error) {
	return &types.Item{}, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetRarities() (*[]types.Rarity, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetItemTypes() (*[]types.ItemType, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetImages() (*[]types.Image, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetAttributes() (*[]types.Attribute, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) GetAttributeGroupings() (*[]types.AttributeGrouping, error) {
	return nil, errors.New("Unsupported")
}

func (p *SqLiteStorageProvider) ItemNameExistsInDb(string) bool {
	return true
}

func (p *SqLiteStorageProvider) UpdateItem(uint64, *types.Item) (*types.Item, error) {
	return &types.Item{}, errors.New("Unsupported")
}

var _ StorageProvider = (*SqLiteStorageProvider)(nil)
