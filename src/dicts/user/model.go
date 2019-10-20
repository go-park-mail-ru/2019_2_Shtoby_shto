package user

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts"
)

const userTableName = "users"

type User struct {
	dicts.BaseInfo
	Login    string                 `json:"login, omitempty" sql:"not null;unique"`
	Password string                 `json:"password,omitempty" sql:"not null"`
	PhotoID  *customType.StringUUID `json:"photo_id,omitempty" sql:"type:uuid"`
	//FirstName string `json:"first_name"`
	//LastName  string `json:"last_name"`
	//Email     string `json:"email"`
}

func (u User) GetTableName() string {
	return userTableName
}

func (u User) IsValid() bool {
	return u.Login != ""
}
