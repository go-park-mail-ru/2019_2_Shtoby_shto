package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const cardGroupsTableName = "card_groups"

//easyjson:json
type CardGroup struct {
	dicts.BaseInfo
	Name    string                `json:"name"`
	BoardID customType.StringUUID `json:"board_id"`
	Cards   []Card                `json:"cards" sql:"-"`
}

func (b CardGroup) GetTableName() string {
	return cardGroupsTableName
}

func (b CardGroup) IsValid() bool {
	return b.Name != ""
}
