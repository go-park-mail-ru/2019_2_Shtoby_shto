package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/task"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption     string                `json:"caption"`
	Priority    int                   `json:"priority"`
	CardUserID  customType.StringUUID `json:"card_user_id"`
	CardGroupID customType.StringUUID `json:"card_group_id"`
	Tasks       []task.Task           `json:"tasks" sql:"-"`
}

func (b Card) GetTableName() string {
	return cardTableName
}

func (b Card) IsValid() bool {
	return b.Caption != "" && b.CardGroupID != ""
}

//easyjson:json
type CardsBoardRequest struct {
	Boards []string `json:"boards"`
}
