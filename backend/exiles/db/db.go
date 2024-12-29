package db

import (
	"sinkzjs.org/m/v2/exiles/types"
)

type Db struct {
	Provider *Provider
}

func NewDatabase(provider *Provider) *Db {
	return &Db{Provider: provider}
}

func (db *Db) GetExiles() (exileArray *[]types.Exile, err error) {
	return (*db.Provider).GetExiles()
}

func (db *Db) GetExile(id uint64) (exile *types.Exile, err error) {
	return (*db.Provider).GetExile(id)
}

func (db *Db) CreateExile(exile *types.Exile) (added_exile *types.Exile, err error) {
	return (*db.Provider).CreateExile(exile)
}

func (db *Db) UpdateExile(id uint64, exile *types.Exile) (updated_exile *types.Exile, err error) {
	return (*db.Provider).UpdateExile(id, exile)
}

func (db *Db) DeleteExile(id uint64) (err error) {
	return (*db.Provider).DeleteExile(id)
}

func (db *Db) ExileNameExistsInDb(name string) bool {
	return (*db.Provider).ExileNameExistsInDb(name)
}
