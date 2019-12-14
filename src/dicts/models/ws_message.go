package models

import "2019_2_Shtoby_shto/src/customType"

//easyjson:json
type WsCardAttachRequest struct {
	UserID customType.StringUUID `json:"user_id"`
	CardID customType.StringUUID `json:"card_id"`
}

//easyjson:json
type RegUser struct {
	UserID customType.StringUUID `json:"user_id"`
}
