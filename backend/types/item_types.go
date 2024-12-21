package types

type ItemType struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Rarity struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	Id  uint64 `json:"id"`
	URL string `json:"url"`
}

type Item struct {
	Id            uint64       `json:"id"`
	Name          string       `json:"name"`
	RequiredLevel int          `json:"required_level"`
	ItemType      ItemType     `json:"item_type"`
	Rarity        Rarity       `json:"rarity"`
	Image         Image        `json:"image"`
	Attributes    *[]Attribute `json:"item_attributes"`
}

type Attribute struct {
	Id                uint64            `json:"id"`
	Name              string            `json:"name"`
	LowValue          int32             `json:"low_value"`
	HighValue         int32             `json:"high_value"`
	AttributeGrouping AttributeGrouping `json:"attribute_grouping"`
}

type AttributeGrouping struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type ItemsResponse struct {
	Items      []Item `json:"items"`
	Message    string `json:"message"`
	HttpStatus int    `json:"request_status"`
}
