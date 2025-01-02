package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"sync"

	"sinkzjs.org/m/v2/items/types"
)

var (
	items              = make(map[uint64]types.Item)
	itemTypes          = make(map[uint64]types.ItemType)
	rarities           = make(map[uint64]types.Rarity)
	images             = make(map[uint64]types.Image)
	attributes         = make(map[uint64]types.Attribute)
	attributeGroupings = make(map[uint64]types.AttributeGrouping)
	loaded             = false
	mu                 sync.RWMutex
)

type FileStorageProvider struct {
	FileName string
	log      slog.Logger
}

func NewFileStorageProvider(fileName string) *FileStorageProvider {
	provider := FileStorageProvider{
		FileName: fileName,
		log:      *slog.Default().With("area", "FileStorageProvider"),
	}
	provider.init()
	return &provider
}

func (p *FileStorageProvider) init() error {
	mu.Lock()
	defer mu.Unlock()

	if loaded {
		return errors.New("data already loaded")
	}

	p.log.Info("Loading data from file " + p.FileName)
	jsonFile, err := os.Open(p.FileName)

	if err != nil {
		p.log.Error(err.Error())
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var itemsArr []types.Item
	if err := json.Unmarshal(byteValue, &itemsArr); err != nil {
		p.log.Error(err.Error())
		return err
	}

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
	p.log.Info("Data loaded")
	return nil
}

func (p *FileStorageProvider) GetItems() (*[]types.Item, error) {
	mu.RLock()
	defer mu.RUnlock()

	if len(items) == 0 {
		return nil, errors.New("no items found")
	}
	itemList := make([]types.Item, 0, len(items))

	for _, item := range items {
		itemList = append(itemList, item)
	}
	return &itemList, nil
}

func (p *FileStorageProvider) GetItemById(id uint64) (*types.Item, error) {
	mu.RLock()
	defer mu.RUnlock()

	item, exists := items[id]
	if !exists {
		return nil, errors.New("item not found")
	}
	return &item, nil
}

func (p *FileStorageProvider) AddItem(item *types.Item) (*types.Item, error) {
	mu.Lock()
	defer mu.Unlock()

	items[item.Id] = *item
	return item, nil
}

func (p *FileStorageProvider) GetRarities() (*[]types.Rarity, error) {
	mu.RLock()
	defer mu.RUnlock()
	if len(rarities) == 0 {
		return nil, errors.New("no rarities found")
	}
	list := make([]types.Rarity, 0, len(rarities))

	for _, item := range rarities {
		list = append(list, item)
	}
	return &list, nil
}

func (p *FileStorageProvider) GetItemTypes() (*[]types.ItemType, error) {
	mu.RLock()
	defer mu.RUnlock()
	if len(itemTypes) == 0 {
		return nil, errors.New("no itemTypes found")
	}
	list := make([]types.ItemType, 0, len(itemTypes))

	for _, value := range itemTypes {
		list = append(list, value)
	}
	return &list, nil
}

func (p *FileStorageProvider) GetImages() (*[]types.Image, error) {
	mu.RLock()
	defer mu.RUnlock()
	if len(images) == 0 {
		return nil, errors.New("no images found")
	}
	list := make([]types.Image, 0, len(images))

	for _, value := range images {
		list = append(list, value)
	}
	return &list, nil
}

func (p *FileStorageProvider) GetAttributes() (*[]types.Attribute, error) {
	mu.RLock()
	defer mu.RUnlock()
	if len(attributes) == 0 {
		return nil, errors.New("no attributes found")
	}
	list := make([]types.Attribute, 0, len(attributes))

	for _, item := range attributes {
		list = append(list, item)
	}
	return &list, nil
}

func (p *FileStorageProvider) GetAttributeGroupings() (*[]types.AttributeGrouping, error) {
	mu.RLock()
	defer mu.RUnlock()
	if len(attributeGroupings) == 0 {
		return nil, errors.New("no attributeGroupings found")
	}
	list := make([]types.AttributeGrouping, 0, len(attributeGroupings))

	for _, item := range attributeGroupings {
		list = append(list, item)
	}
	return &list, nil
}

func (p *FileStorageProvider) ItemNameExistsInDb(name string) bool {
	mu.RLock()
	defer mu.RUnlock()

	for _, item := range items {
		if item.Name == name {
			return true
		}
	}
	return false
}

func (p *FileStorageProvider) UpdateItem(id uint64, item *types.Item) (*types.Item, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := items[id]; !exists {
		return &types.Item{}, errors.New("item not found")
	}
	item.Id = id
	items[id] = *item
	return item, nil
}

var _ StorageProvider = (*FileStorageProvider)(nil)
