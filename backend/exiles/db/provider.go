package db

import (
	"sinkzjs.org/m/v2/exiles/types"
)

type Provider interface {
	GetExiles() (*[]types.Exile, error)
	GetExile(id uint64) (*types.Exile, error)
	CreateExile(exile *types.Exile) (*types.Exile, error)
	UpdateExile(id uint64, exile *types.Exile) (*types.Exile, error)
	DeleteExile(id uint64) error
	ExileNameExistsInDb(name string) bool
}
