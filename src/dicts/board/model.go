package photo

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const boardTableName = "board"

type Board struct {
	dicts.BaseInfo
}

func (p Board) GetTableName() string {
	return boardTableName
}
