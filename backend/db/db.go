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

func LoadData() {
	// Load data from JSON file
	var jsonFile, err = os.Open("data/data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var itemsArr []types.Item
	json.Unmarshal(byteValue, &itemsArr)

	for _, item := range itemsArr {
		items[item.Id] = item
	}
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
	if ok {
		return item, fmt.Errorf("item already exists in system: %s [%d]", item.Name, item.Id)
	}

	id := uint64(len(items) + 1)
	item.Id = id

	items[id] = item
	return item, nil
}
