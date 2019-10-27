package сardUsers

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const cardUsersTableName = "cardUsers"

//easyjson:json
type CardUsers struct {
	dicts.BaseInfo
	CardID string `json:"card_id"`
	UserID string `json:"user_id"`
}

func (b CardUsers) GetTableName() string {
	return cardUsersTableName
}

func (b CardUsers) IsValid() bool {
	return b.CardID != "" && b.UserID != ""
}
