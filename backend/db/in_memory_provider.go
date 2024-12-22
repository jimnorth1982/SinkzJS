package db

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"sinkzjs.org/m/v2/types"
)

var (
	items              = make(map[uint64]types.Item)
	itemTypes          = make(map[uint64]types.ItemType)
	rarities           = make(map[uint64]types.Rarity)
	images             = make(map[uint64]types.Image)
	attributes         = make(map[uint64]types.Attribute)
	attributeGroupings = make(map[uint64]types.AttributeGrouping)
	loaded             = false
	mu                 sync.Mutex
)

type InMemoryProvider struct{}

func NewInMemoryProvider() *InMemoryProvider {
	provider := InMemoryProvider{}
	provider.loadData()
	return &provider
}

func (p *InMemoryProvider) loadData() error {
	mu.Lock()
	defer mu.Unlock()

	log.Println("InMemoryProvider: Loading data from file")
	if loaded {
		return nil
	}

	jsonFile, err := os.Open("data/item_data.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var itemsArr []types.Item
	if err := json.Unmarshal(byteValue, &itemsArr); err != nil {
		log.Println(err)
		return err
	}

	log.Println("InMemoryProvider: Parsing data")
	for _, item := range itemsArr {
		items[item.Id] = item
		rarities[item.Rarity.Id] = item.Rarity
		itemTypes[item.ItemType.Id] = item.ItemType
		images[item.Image.Id] = item.Image
		if item.Attributes != nil {
			for _, attribute := range *item.Attributes {
				attributes[attribute.Id] = attribute
				attributeGroupings[attribute.AttributeGrouping.Id] = attribute.AttributeGrouping
			}
		}
	}
	loaded = true
	log.Println("InMemoryProvider: Data loaded successfully: ", len(items), " items")
	return nil
}

func (p *InMemoryProvider) GetItems() ([]types.Item, error) {
	mu.Lock()
	defer mu.Unlock()

	if len(items) == 0 {
		return nil, errors.New("no items found")
	}
	itemList := make([]types.Item, 0, len(items))

	for _, item := range items {
		itemList = append(itemList, item)
	}
	return itemList, nil
}

func (p *InMemoryProvider) GetItemById(id uint64) (types.Item, error) {
	mu.Lock()
	defer mu.Unlock()

	item, exists := items[id]
	if !exists {
		return types.Item{}, errors.New("item not found")
	}
	return item, nil
}

func (p *InMemoryProvider) AddItem(item types.Item) (types.Item, error) {
	mu.Lock()
	defer mu.Unlock()

	items[item.Id] = item
	return item, nil
}

func (p *InMemoryProvider) GetRarities() (map[uint64]types.Rarity, error) {
	mu.Lock()
	defer mu.Unlock()

	return rarities, nil
}

func (p *InMemoryProvider) GetItemTypes() (map[uint64]types.ItemType, error) {
	mu.Lock()
	defer mu.Unlock()

	return itemTypes, nil
}

func (p *InMemoryProvider) GetImages() (map[uint64]types.Image, error) {
	mu.Lock()
	defer mu.Unlock()

	return images, nil
}

func (p *InMemoryProvider) GetAttributeGroupings() (attributeGroupingList map[uint64]types.AttributeGrouping, err error) {
	if len(attributeGroupings) == 0 {
		return nil, errors.New("no attribute groupings found")
	}

	return attributeGroupings, nil
}

func (p *InMemoryProvider) GetAttributes() (map[uint64]types.Attribute, error) {
	mu.Lock()
	defer mu.Unlock()

	return attributes, nil
}

func (p *InMemoryProvider) ItemNameExistsInDb(name string) bool {
	mu.Lock()
	defer mu.Unlock()

	for _, item := range items {
		if item.Name == name {
			return true
		}
	}
	return false
}

var _ Provider = (*InMemoryProvider)(nil)
