package models

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const commentTableName = "comments"

//easyjson:json
type Comment struct {
	dicts.BaseInfo
	Text   string `json:"text"`
	CardID string `json:"card_id"`
	UserID string `json:"user_id"`
	//Done   bool   `json:"done"`
}

func (c Comment) GetTableName() string {
	return commentTableName
}

func (c Comment) IsValid() bool {
	return c.Text != ""
}
