package db

import (
	"sync"
	"testing"

	"sinkzjs.org/m/v2/items/types"
)

var provider = NewInMemoryProvider("data/item_data.json")

func TestConcurrentAccess(t *testing.T) {
	var wg sync.WaitGroup

	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_, err := provider.GetItems()
			if err != nil {
				t.Errorf("Failed to get items: %v", err)
			}
		}()
	}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id uint64) {
			defer wg.Done()
			item := types.Item{
				Id:   id,
				Name: "Test Item",
				ItemType: types.ItemType{
					Id:   1,
					Name: "Test Type",
				},
				Rarity: types.Rarity{
					Id:   1,
					Name: "Common",
				},
				Image: types.Image{
					Id:  1,
					URL: "http://example.com/image.png",
				},
				Attributes: &[]types.Attribute{
					{
						Id:        1,
						Name:      "Power",
						LowValue:  10,
						HighValue: 20,
						AttributeGrouping: types.AttributeGrouping{
							Id:   1,
							Name: "Magic",
						},
					},
				},
			}
			_, err := provider.AddItem(item)
			if err != nil {
				t.Errorf("Failed to add item: %v", err)
			}
		}(uint64(i + 1))
	}

	wg.Wait()
}

func TestItemNameExistsInDb(t *testing.T) {
	item := types.Item{
		Id:   1,
		Name: "Unique Item",
		ItemType: types.ItemType{
			Id:   1,
			Name: "Test Type",
		},
		Rarity: types.Rarity{
			Id:   1,
			Name: "Common",
		},
		Image: types.Image{
			Id:  1,
			URL: "http://example.com/image.png",
		},
		Attributes: &[]types.Attribute{
			{
				Id:        1,
				Name:      "Power",
				LowValue:  10,
				HighValue: 20,
				AttributeGrouping: types.AttributeGrouping{
					Id:   1,
					Name: "Magic",
				},
			},
		},
	}
	_, err := provider.AddItem(item)
	if err != nil {
		t.Fatalf("Failed to add item: %v", err)
	}

	exists := provider.ItemNameExistsInDb("Unique Item")
	if !exists {
		t.Errorf("Expected item name to exist in the database")
	}

	exists = provider.ItemNameExistsInDb("Non-Existent Item")
	if exists {
		t.Errorf("Expected item name to not exist in the database")
	}
}