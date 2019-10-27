package card

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const cardTableName = "cards"

//easyjson:json
type Card struct {
	dicts.BaseInfo
	Caption     string                `json:"caption"`
	Priority    int                   `json:"priority"`
	BoardID     customType.StringUUID `json:"board_id"`
	CardUserID  customType.StringUUID `json:"card_user_id"`
	CardGroupID customType.StringUUID `json:"card_group_id"`
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
