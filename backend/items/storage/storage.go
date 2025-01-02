package storage

import "sinkzjs.org/m/v2/items/types"

type Storage struct {
	Provider StorageProvider
}

func NewStorage(provider StorageProvider) *Storage {
	return &Storage{Provider: provider}
}

func (db *Storage) GetItems() (itemArray *[]types.Item, err error) {
	return db.Provider.GetItems()
}

func (db *Storage) GetItemById(id uint64) (item *types.Item, err error) {
	return db.Provider.GetItemById(id)
}

func (db *Storage) AddItem(item *types.Item) (added_item *types.Item, err error) {
	return db.Provider.AddItem(item)
}

func (db *Storage) GetRarities() (rarityList *[]types.Rarity, err error) {
	return db.Provider.GetRarities()
}

func (db *Storage) GetItemTypes() (itemTypeList *[]types.ItemType, err error) {
	return db.Provider.GetItemTypes()
}

func (db *Storage) GetImages() (imageList *[]types.Image, err error) {
	return db.Provider.GetImages()
}

func (db *Storage) GetAttributes() (attributeList *[]types.Attribute, err error) {
	return db.Provider.GetAttributes()
}

func (db *Storage) GetAttributeGroupings() (attributeGroupingList *[]types.AttributeGrouping, err error) {
	return db.Provider.GetAttributeGroupings()
}

func (db *Storage) ItemNameExistsInDb(name string) bool {
	return db.Provider.ItemNameExistsInDb(name)
}
