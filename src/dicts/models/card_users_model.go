package models

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const cardUsersTableName = "card_users"

//easyjson:json
type CardUsers struct {
	dicts.BaseInfo
	CardID customType.StringUUID `json:"card_id"`
	UserID customType.StringUUID `json:"user_id"`
}

func (c CardUsers) GetTableName() string {
	return cardUsersTableName
}

func (c CardUsers) IsValid() bool {
	return c.CardID != "" && c.UserID != ""
}

//easyjson:json
type CardsUserRequest struct {
	Users []string `json:"users"`
}
