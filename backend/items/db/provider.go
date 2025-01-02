package db

import "sinkzjs.org/m/v2/items/types"

type Provider interface {
	GetItems() (*[]types.Item, error)
	GetItemById(uint64) (*types.Item, error)
	AddItem(*types.Item) (*types.Item, error)
	GetRarities() (*[]types.Rarity, error)
	GetItemTypes() (*[]types.ItemType, error)
	GetImages() (*[]types.Image, error)
	GetAttributes() (*[]types.Attribute, error)
	GetAttributeGroupings() (*[]types.AttributeGrouping, error)
	ItemNameExistsInDb(string) bool
	UpdateItem(uint64, *types.Item) (*types.Item, error)
}
