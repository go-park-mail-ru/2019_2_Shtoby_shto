package user

import "2019_2_Shtoby_shto/src/dicts"

const userTableName = "users"

type User struct {
	dicts.BaseInfo
	Login    string `json:"login" sql:"not null;unique"`
	Password string `json:"password" sql:"not null"`
	//FirstName string `json:"first_name"`
	//LastName  string `json:"last_name"`
	//Email     string `json:"email"`
}

func (u User) GetTableName() string {
	return userTableName
}
