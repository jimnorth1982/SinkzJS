package types

import itemTypes "sinkzjs.org/m/v2/items/types"

type Exile struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
	HelmetId uint64 `json:"helmet_id"`
	ArmorId  uint64 `json:"armor_id"`
	WeaponId uint64 `json:"weapon_id"`
	ShieldId uint64 `json:"shield_id"`
	BootsId  uint64 `json:"boots_id"`
}

type HydratedExile struct {
	ID     uint64         `json:"id"`
	Name   string         `json:"name"`
	Level  uint64         `json:"level"`
	Helmet itemTypes.Item `json:"helmet"`
	Armor  itemTypes.Item `json:"armor"`
	Weapon itemTypes.Item `json:"weapon"`
	Shield itemTypes.Item `json:"shield"`
	Boots  itemTypes.Item `json:"boots"`
}

type ExilesResponse struct {
	Message    string  `json:"message"`
	HttpStatus int     `json:"http_status"`
	Exiles     []Exile `json:"exiles"`
}
