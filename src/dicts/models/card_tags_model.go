package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const cardTagsTableName = "card_tags"

//easyjson:json
type CardTags struct {
	dicts.BaseInfo
	CardID customType.StringUUID `json:"card_id"`
	TagID  customType.StringUUID `json:"tag_id"`
}

func (c CardTags) GetTableName() string {
	return cardTagsTableName
}

func (c CardTags) IsValid() bool {
	return c.CardID != "" && c.TagID != ""
}

//easyjson:json
type CardTagsRequest struct {
	Cards []string `json:"cards"`
}
