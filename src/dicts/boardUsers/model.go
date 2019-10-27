package boardUsers

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const boardUsersTableName = "boardUsers"

//easyjson:json
type BoardUsers struct {
	dicts.BaseInfo
	BoardID string `json:"board_id"`
	UserID  string `json:"user_id"`
}

func (b BoardUsers) GetTableName() string {
	return boardUsersTableName
}

func (b BoardUsers) IsValid() bool {
	return b.BoardID != "" && b.UserID != ""
}
