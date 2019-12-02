package models

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const checkListTableName = "check_lists"

//easyjson:json
type CheckList struct {
	dicts.BaseInfo
	Text   string `json:"text"`
	CardID string `json:"card_id"`
	UserID string `json:"user_id"`
	Done   bool   `json:"done"`
}

func (c CheckList) GetTableName() string {
	return checkListTableName
}

func (c CheckList) IsValid() bool {
	return c.Text != ""
}
