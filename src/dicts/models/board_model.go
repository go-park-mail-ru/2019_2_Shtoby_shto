package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const boardTableName = "boards"

//easyjson:json
type Board struct {
	dicts.BaseInfo
	Name         string                `json:"name"`
	ShortURL     string                `json:"short_url"`
	Users        []string              `json:"users" sql:"-"`
	BoardUsersID customType.StringUUID `json:"-"`
	CardGroups   []CardGroup           `json:"card_groups" sql:"-"`
}

func (b Board) GetTableName() string {
	return boardTableName
}

func (b Board) IsValid() bool {
	return b.Name != ""
}

//easyjson:json
type ShortURL struct {
	ShortURL string `json:"short_url"`
}

//easyjson:json
type EventMessage struct {
	Event string `json:"event"`
}
