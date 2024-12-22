package db

import "sinkzjs.org/m/v2/types"

type Db struct {
	provider Provider
}

func NewDatabase(provider Provider) *Db {
	return &Db{provider: provider}
}

func (db *Db) GetItems() (itemArray []types.Item, err error) {
	return db.provider.GetItems()
}

func (db *Db) GetItemById(id uint64) (item types.Item, err error) {
	return db.provider.GetItemById(id)
}

func (db *Db) AddItem(item types.Item) (added_item types.Item, err error) {
	return db.provider.AddItem(item)
}

func (db *Db) GetRarities() (rarityList map[uint64]types.Rarity, err error) {
	return db.provider.GetRarities()
}

func (db *Db) GetItemTypes() (itemTypeList map[uint64]types.ItemType, err error) {
	return db.provider.GetItemTypes()
}

func (db *Db) GetImages() (imageList map[uint64]types.Image, err error) {
	return db.provider.GetImages()
}

func (db *Db) GetAttributes() (attributeList map[uint64]types.Attribute, err error) {
	return db.provider.GetAttributes()
}

func (db *Db) GetAttributeGroupings() (attributeGroupingList map[uint64]types.AttributeGrouping, err error) {
	return db.provider.GetAttributeGroupings()
}

func (db *Db) ItemNameExistsInDb(name string) bool {
	return db.provider.ItemNameExistsInDb(name)
}
