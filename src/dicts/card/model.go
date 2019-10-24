package card

import (
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/user"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption string `json:"caption"`
	// Tasks []task.Task `json:"tasks"`
	UsersAssign []user.User `json:"users" sql:"-"`
}

func (b Card) GetTableName() string {
	return cardTableName
}

func (b Card) IsValid() bool {
	return b.Caption != ""
}
