package controller

import (
	"GoAuth/dto"
	"GoAuth/services"
	util "GoAuth/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		util.WriteJsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	util.WriteJsonSuccessResponse(w, "User logged in Successfully", jwtToken)
	
}

func (uc *UserController) FetchUserbyId(w http.ResponseWriter, r *http.Request){
	fmt.Println("fetching user by Id")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64) 
	if err != nil{
		util.WriteJsonErrorResponse(w, fmt.Errorf("invalid ID: %v", err), http.StatusBadRequest)
	}

	iderr := uc.userService.FindUserById(id)
	if iderr != nil{
		util.WriteJsonErrorResponse(w, err, http.StatusNotFound)
		return
	}
	w.Write([]byte("fetch by id endpoint"))
}

func (uc *UserController) FetchAllUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("fetching all users")
	uc.userService.GetAllUsers()
	w.Write([]byte("fetching all users endpoint"))
}
