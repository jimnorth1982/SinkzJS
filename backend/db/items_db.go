package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"sinkzjs.org/m/v2/types"
)

var items = map[uint64]types.Item{}
var itemTypes = map[uint64]types.ItemType{}
var rarities = map[uint64]types.Rarity{}
var images = map[uint64]types.Image{}
var attributes = map[uint64]types.Attribute{}
var attributeGroupings = map[uint64]types.AttributeGrouping{}
var loaded = false

func LoadData() {
	if loaded {
		return
	}
	var jsonFile, err = os.Open("data/item_data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var itemsArr []types.Item
	json.Unmarshal(byteValue, &itemsArr)

	for _, item := range itemsArr {
		items[item.Id] = item
		rarities[item.Rarity.Id] = item.Rarity
		itemTypes[item.ItemType.Id] = item.ItemType
		images[item.Image.Id] = item.Image
		for _, attribute := range *item.Attributes {
			attributes[attribute.Id] = attribute
			attributeGroupings[attribute.AttributeGrouping.Id] = attribute.AttributeGrouping
		}
	}
	loaded = true
}

func GetItems() (itemList map[uint64]types.Item, err error) {
	if len(items) == 0 {
		return nil, errors.New("no items found")
	}

	return items, nil
}

func GetItemById(id uint64) (fetched_item *types.Item, err error) {
	item, ok := items[id]
	if !ok || item.Id == 0 {
		return nil, fmt.Errorf("item not found: %d", id)
	}
	return &item, nil
}

func AddItem(item types.Item) (added_item types.Item, err error) {
	item, ok := items[item.Id]
	if ok || itemNameExists(item.Name) {
		return item, fmt.Errorf("item already exists in system: %s [%d]", item.Name, item.Id)
	}

	id := uint64(len(items) + 1)
	item.Id = id

	items[id] = item
	return item, nil
}

func GetRarities() (rarityList map[uint64]types.Rarity, err error) {
	if len(rarities) == 0 {
		return nil, errors.New("no rarities found")
	}

	return rarities, nil
}

func GetItemTypes() (itemTypeList map[uint64]types.ItemType, err error) {
	if len(itemTypes) == 0 {
		return nil, errors.New("no item types found")
	}

	return itemTypes, nil
}

func GetImages() (imageList map[uint64]types.Image, err error) {
	if len(images) == 0 {
		return nil, errors.New("no images found")
	}

	return images, nil
}

func GetAttributes() (attributeList map[uint64]types.Attribute, err error) {
	if len(attributes) == 0 {
		return nil, errors.New("no attributes found")
	}

	return attributes, nil
}

func GetAttributeGroupings() (attributeGroupingList map[uint64]types.AttributeGrouping, err error) {
	if len(attributeGroupings) == 0 {
		return nil, errors.New("no attribute groupings found")
	}

	return attributeGroupings, nil
}

func itemNameExists(name string) bool {
	for _, item := range items {
		if item.Name == name {
			return true
		}
	}
	return false
}
