package card

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption     string `json:"caption"`
	Priority    int    `json:"priority"`
	BoardID     string `json:"board_id"`
	CardGroupID string `json:"card_group_id"`
}

func (b Card) GetTableName() string {
	return cardTableName
}

func (b Card) IsValid() bool {
	return b.Caption != ""
}
