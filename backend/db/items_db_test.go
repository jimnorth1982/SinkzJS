package db

import (
	"testing"

	"sinkzjs.org/m/v2/types"
)

func setup() {
	items = map[uint64]types.Item{
		1: {Id: 1, Name: "Item1", Rarity: types.Rarity{Id: 1, Name: "Common"}, ItemType: types.ItemType{Id: 1, Name: "Type1"}, Image: types.Image{Id: 1, URL: "http://example.com/image1"}, Attributes: &[]types.Attribute{{Id: 1, Name: "Attr1", AttributeGrouping: types.AttributeGrouping{Id: 1, Name: "Group1"}}}},
		2: {Id: 2, Name: "Item2", Rarity: types.Rarity{Id: 1, Name: "Common"}, ItemType: types.ItemType{Id: 1, Name: "Type1"}, Image: types.Image{Id: 1, URL: "http://example.com/image1"}, Attributes: &[]types.Attribute{{Id: 1, Name: "Attr1", AttributeGrouping: types.AttributeGrouping{Id: 1, Name: "Group1"}}}},
	}
	rarities = map[uint64]types.Rarity{
		1: {Id: 1, Name: "Common"},
	}
	itemTypes = map[uint64]types.ItemType{
		1: {Id: 1, Name: "Type1"},
	}
	images = map[uint64]types.Image{
		1: {Id: 1, URL: "http://example.com/image1"},
	}
	attributes = map[uint64]types.Attribute{
		1: {Id: 1, Name: "Attr1", AttributeGrouping: types.AttributeGrouping{Id: 1, Name: "Group1"}},
	}
	attributeGroupings = map[uint64]types.AttributeGrouping{
		1: {Id: 1, Name: "Group1"},
	}
	loaded = true
}

func TestGetItems(t *testing.T) {
	setup()
	itemList, err := GetItems()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(itemList) != 2 {
		t.Fatalf("expected 1 item, got %d", len(itemList))
	}
}

func TestGetItemById(t *testing.T) {
	setup()
	item, err := GetItemById(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if item.Id != 1 {
		t.Fatalf("expected item ID 1, got %d", item.Id)
	}
}

func TestAddItem(t *testing.T) {
	setup()
	newItem := types.Item{Id: 3, Name: "Item3", Rarity: types.Rarity{Id: 2, Name: "Rare"}, ItemType: types.ItemType{Id: 2, Name: "Type2"}, Image: types.Image{Id: 2, URL: "http://example.com/image2"}, Attributes: &[]types.Attribute{{Id: 2, Name: "Attr2", AttributeGrouping: types.AttributeGrouping{Id: 2, Name: "Group2"}}}}
	addedItem, err := AddItem(newItem)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if addedItem.Id != 3 {
		t.Fatalf("expected item ID 3, got %d", addedItem.Id)
	}
}

func TestGetRarities(t *testing.T) {
	setup()
	rarityList, err := GetRarities()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(rarityList) != 1 {
		t.Fatalf("expected 1 rarity, got %d", len(rarityList))
	}
}

func TestGetItemTypes(t *testing.T) {
	setup()
	itemTypeList, err := GetItemTypes()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(itemTypeList) != 1 {
		t.Fatalf("expected 1 item type, got %d", len(itemTypeList))
	}
}

func TestGetImages(t *testing.T) {
	setup()
	imageList, err := GetImages()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(imageList) != 1 {
		t.Fatalf("expected 1 image, got %d", len(imageList))
	}
}

func TestGetAttributes(t *testing.T) {
	setup()
	attributeList, err := GetAttributes()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(attributeList) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(attributeList))
	}
}

func TestGetAttributeGroupings(t *testing.T) {
	setup()
	attributeGroupingList, err := GetAttributeGroupings()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(attributeGroupingList) != 1 {
		t.Fatalf("expected 1 attribute grouping, got %d", len(attributeGroupingList))
	}
}

func TestItemNameExists(t *testing.T) {
	setup()
	exists := itemNameExists("Item1")
	if !exists {
		t.Fatalf("expected item name to exist")
	}
	exists = itemNameExists("NonExistentItem")
	if exists {
		t.Fatalf("expected item name to not exist")
	}
}
