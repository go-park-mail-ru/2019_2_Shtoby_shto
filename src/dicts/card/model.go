package card

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption string `json:"caption"`
	TaskID  string `json:"task_id"`
}

func (b Card) GetTableName() string {
	return cardTableName
}

func (b Card) IsValid() bool {
	return b.Caption != ""
}
