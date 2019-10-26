package board

import (
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/user"
)

const boardTableName = "boards"

//easyjson:json
type Board struct {
	dicts.BaseInfo
	Name 	string `json:"name"`
	Users 	[]user.User `json:"users" sql:"-"`
}

func (b Board) GetTableName() string {
	return boardTableName
}

func (b Board) IsValid() bool {
	return b.Name != ""
}
