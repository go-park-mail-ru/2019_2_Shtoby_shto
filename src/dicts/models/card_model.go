package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption     string                `json:"caption"`
	Text        string                `json:"text"`
	Priority    int                   `json:"priority"`
	FileID      customType.StringUUID `json:"file_id"`
	CardUserID  customType.StringUUID `json:"-"`
	CardGroupID customType.StringUUID `json:"card_group_id"`
	File        customType.StringUUID `json:"file"`
	Comments    []Comment             `json:"comments" sql:"-"`
	Tags        []Tag                 `json:"tags" sql:"-"`
	Users       []string              `json:"users" sql:"-"`
	CheckLists  []CheckList           `json:"check_lists" sql:"-"`
}

func (c Card) GetTableName() string {
	return cardTableName
}

func (c Card) IsValid() bool {
	return c.Caption != "" && c.CardGroupID != ""
}

//easyjson:json
type CardsBoardRequest struct {
	Boards []string `json:"boards"`
}
