package boardUsers

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const boardUsersTableName = "board_users"

//easyjson:json
type BoardUsers struct {
	dicts.BaseInfo
	BoardID customType.StringUUID `json:"board_id"`
	UserID  customType.StringUUID `json:"user_id"`
}

func (b BoardUsers) GetTableName() string {
	return boardUsersTableName
}

func (b BoardUsers) IsValid() bool {
	return b.BoardID != "" && b.UserID != ""
}
