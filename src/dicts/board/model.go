package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
)

const boardTableName = "boards"

//easyjson:json
type Board struct {
	dicts.BaseInfo
	Name         string                `json:"name"`
	BoardUsersID customType.StringUUID `json:"board_users_id"`
	CardGroups   []cardGroup.CardGroup `json:"card_groups" sql:"-"`
}

func (b Board) GetTableName() string {
	return boardTableName
}

func (b Board) IsValid() bool {
	return b.Name != ""
}
