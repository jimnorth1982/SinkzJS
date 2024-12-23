package db

import "sinkzjs.org/m/v2/types"

type Provider interface {
	GetItems() ([]types.Item, error)
	GetItemById(uint64) (types.Item, error)
	AddItem(types.Item) (types.Item, error)
	GetRarities() (map[uint64]types.Rarity, error)
	GetItemTypes() (map[uint64]types.ItemType, error)
	GetImages() (map[uint64]types.Image, error)
	GetAttributes() (map[uint64]types.Attribute, error)
	GetAttributeGroupings() (map[uint64]types.AttributeGrouping, error)
	ItemNameExistsInDb(string) bool
	Init() error
}
