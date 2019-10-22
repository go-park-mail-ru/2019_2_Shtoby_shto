package board

import (
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/user"
)

const boardTableName = "board"

//easyjson:json
type Board struct {
	dicts.BaseInfo
	Caption string `json:"caption"`
	// Tasks []task.Task `json:"tasks"`
	Users []user.User `json:"users" sql:"-"`
}

func (b Board) GetTableName() string {
	return boardTableName
}

func (b Board) IsValid() bool {
	return b.Caption != ""
}
