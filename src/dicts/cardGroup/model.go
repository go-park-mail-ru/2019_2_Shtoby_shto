package cardGroup

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const cardGroupsTableName = "card_groups"

//easyjson:json
type CardGroup struct {
	dicts.BaseInfo
	Name string `json:"name"`
}

func (b CardGroup) GetTableName() string {
	return cardGroupsTableName
}

func (b CardGroup) IsValid() bool {
	return b.Name != ""
}
