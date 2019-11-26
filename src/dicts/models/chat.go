package models

import (
	"2019_2_Shtoby_shto/src/dicts"
	"time"
)

const messageTableName = "message"

//easyjson:json
type Message struct {
	dicts.BaseInfo
	Author string `json:"authorId"`
	//Receiver string   `json:"toId"`
	Created time.Time `json:"created"`
	//Edited   bool      `json:"isEdited"`
	Text string `json:"text"`
}

func (m Message) GetTableName() string {
	return messageTableName
}

func (m Message) IsValid() bool {
	return m.Text != ""
}
