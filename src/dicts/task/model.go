package task

import (
	"2019_2_Shtoby_shto/src/dicts"
)

const taskTableName = "tasks"

//easyjson:json
type Task struct {
	dicts.BaseInfo
	Text   string `json:"text"`
	CardID string `json:"card_id"`
	Done   bool   `json:"done"`
}

func (b Task) GetTableName() string {
	return taskTableName
}

func (b Task) IsValid() bool {
	return b.Text != ""
}
