package controller

import (
	"GoAuth/dto"
	"GoAuth/services"
	util "GoAuth/utils"
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

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("Login in the user")
	
	payload := r.Context().Value("payload").(dto.LoginUserRequestDto)
	fmt.Println("Payload received in controller layer:", payload)

	jwtToken, err := uc.userService.LoginUser(&payload)
	if err != nil{
		util.WriteJsonErrorResponse(w, "User error logging in", "unable to create login token")
		return
	}
	util.WriteJsonSuccessResponse(w, "User logged in Successfully", jwtToken)
	
}

func (uc *UserController) FetchUserbyId(w http.ResponseWriter, r *http.Request){
	fmt.Println("fetching user by Id")
	uc.userService.UserById()
	w.Write([]byte("fetch by id endpoint"))
}

func (uc *UserController) FetchAllUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("fetching all users")
	uc.userService.GetAllUsers()
	w.Write([]byte("fetching all users endpoint"))
}
