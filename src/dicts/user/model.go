package user

import "2019_2_Shtoby_shto/src/dicts"

const userTableName = "users"

type User struct {
	dicts.BaseInfo
	Login    string `json:"login" sql:"login"`
	Password string `json:"password" sql:"password"`
	//FirstName string `json:"first_name"`
	//LastName  string `json:"last_name"`
	//Email     string `json:"email"`
}

func (u User) GetTableName() string {
	return userTableName
}
