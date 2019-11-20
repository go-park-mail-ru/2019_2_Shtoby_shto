package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"time"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption     string                `json:"caption"`
	Text        string                `json:"text"`
	Priority    int                   `json:"priority"`
	CardUserID  customType.StringUUID `json:"-"`
	CardGroupID customType.StringUUID `json:"card_group_id"`
	Comments    []Comment             `json:"comments" sql:"-"`
	Tags        []Tag                 `json:"tags" sql:"-"`
	Users       []string              `json:"users" sql:"-"`
	Deadline    time.Time             `json:"deadline"`
}

func (b Card) GetTableName() string {
	return cardTableName
}

func (b Card) IsValid() bool {
	return b.Caption != "" && b.CardGroupID != ""
}

//easyjson:json
type CardsBoardRequest struct {
	Boards []string `json:"boards"`
}
