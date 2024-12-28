package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"sinkzjs.org/m/v2/items/types"
)

type SqLiteDb struct {
	db *sql.DB
}

type SqLiteProvider struct {
	DBFileName    string
	TimeoutMillis int
	TTLMillis     int
	sqlitedb      SqLiteDb
}

func NewSqlite3Provider(DBFileName string, TimeoutMillis int, TTLMillis int) *SqLiteProvider {
	conn := Connect(DBFileName)
	provider := &SqLiteProvider{DBFileName, TimeoutMillis, TTLMillis, *conn}
	provider.init()
	return provider
}

func (p *SqLiteProvider) init() {
	log.Print("Init: SqLiteProvider")
}

func Connect(dbpath string) *SqLiteDb {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatalf("DB Handle Error: %s", err.Error())
	}
	return &SqLiteDb{db}
}

func (p *SqLiteProvider) GetItems() ([]types.Item, error) {
	return nil, nil
}

func (p *SqLiteProvider) GetItemById(uint64) (types.Item, error) {
	return types.Item{}, nil
}

func (p *SqLiteProvider) AddItem(types.Item) (types.Item, error) {
	return types.Item{}, nil
}

func (p *SqLiteProvider) GetRarities() (map[uint64]types.Rarity, error) {
	return nil, nil
}

func (p *SqLiteProvider) GetItemTypes() (map[uint64]types.ItemType, error) {
	return nil, nil
}

func (p *SqLiteProvider) GetImages() (map[uint64]types.Image, error) {
	return nil, nil
}

func (p *SqLiteProvider) GetAttributes() (map[uint64]types.Attribute, error) {
	return nil, nil
}

func (p *SqLiteProvider) GetAttributeGroupings() (map[uint64]types.AttributeGrouping, error) {
	return nil, nil
}

func (p *SqLiteProvider) ItemNameExistsInDb(string) bool {
	return true
}

func (p *SqLiteProvider) UpdateItem(uint64, types.Item) (types.Item, error) {
	return types.Item{}, nil
}

var _ Provider = (*SqLiteProvider)(nil)
