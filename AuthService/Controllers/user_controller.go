package controller

import (
	"GoAuth/services"
	"fmt"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController)RegisterUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("Controller layer")
	uc.userService.CreateUser()
	w.Write([]byte("User Registeration endpoint"))
}