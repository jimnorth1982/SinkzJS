package db

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"sinkzjs.org/m/v2/exiles/types"
)

var (
	Exiles = make(map[uint64]types.Exile)
	loaded = false
	mu     sync.RWMutex
)

type InMemoryProvider struct {
	FileName string
}

func NewInMemoryProvider(fileName string) *InMemoryProvider {
	provider := &InMemoryProvider{FileName: fileName}
	provider.init()
	return provider
}

func (p *InMemoryProvider) init() error {
	mu.Lock()
	defer mu.Unlock()

	if loaded {
		return errors.New("data already loaded")
	}

	log.Println("Loading data from file:", p.FileName)
	jsonFile, err := os.Open(p.FileName)

	if err != nil {
		log.Fatal(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var exilesArray []types.Exile
	if err := json.Unmarshal(byteValue, &exilesArray); err != nil {
		log.Fatal(err)
		return err
	}

	for _, exile := range exilesArray {
		Exiles[exile.ID] = exile
	}

	loaded = true
	log.Println("Data loaded")
	return nil
}

func (p *InMemoryProvider) GetExiles() ([]types.Exile, error) {
	var exiles []types.Exile
	for _, exile := range Exiles {
		exiles = append(exiles, exile)
	}
	return exiles, nil
}

func (p *InMemoryProvider) GetExile(id uint64) (types.Exile, error) {
	return Exiles[id], nil
}

func (p *InMemoryProvider) CreateExile(exile types.Exile) (types.Exile, error) {
	Exiles[exile.ID] = exile
	return exile, nil
}

func (p *InMemoryProvider) UpdateExile(id uint64, exile types.Exile) (types.Exile, error) {
	Exiles[id] = exile
	return exile, nil
}

func (p *InMemoryProvider) DeleteExile(id uint64) error {
	delete(Exiles, id)
	return nil
}

func (p *InMemoryProvider) ExileNameExistsInDb(name string) bool {
	for _, exile := range Exiles {
		if exile.Name == name {
			return true
		}
	}
	return false
}

var _ Provider = (*InMemoryProvider)(nil)
