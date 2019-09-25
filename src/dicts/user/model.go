package user

type User struct {
	Id       string `json:"id" sql:"id"`
	Login    string `json:"login" sql:"login"`
	Password string `json:"password" sql:"password"`
	//FirstName string `json:"first_name"`
	//LastName  string `json:"last_name"`
	//Email     string `json:"email"`
}
