package task

import (
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/dicts/user"
)

const taskTableName = "tasks"

//easyjson:json
type Task struct {
	dicts.BaseInfo
	Caption string `json:"caption"`
	// Tasks []task.Task `json:"tasks"`
	UsersAssign []user.User `json:"users" sql:"-"`
}

func (b Task) GetTableName() string {
	return taskTableName
}

func (b Task) IsValid() bool {
	return b.Caption != ""
}
