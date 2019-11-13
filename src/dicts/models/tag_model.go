package models

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const tagTableName = "tags"

//easyjson:json
type Tag struct {
	dicts.BaseInfo
	Text string `json:"text"`
}

func (b Tag) GetTableName() string {
	return tagTableName
}

func (b Tag) IsValid() bool {
	return b.Text != ""
}
