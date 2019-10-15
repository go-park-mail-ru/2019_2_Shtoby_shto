package user

import (
	transport "2019_2_Shtoby_shto/src/handle"
	"net/http"
)

type UserHandler struct {
	userService HandlerUserService
	transport.Handler
}

func NewUserHandler(userService HandlerUserService) {
	//handler := UserHandler{
	//	userService: userService,
	//}

	return
}

func (u UserHandler) Get(w http.ResponseWriter, req *http.Request) {

}
