package router

import (
	controller "GoAuth/Controllers"
	middleware "GoAuth/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controller.UserController
}

func NewUserRouter(_userController *controller.UserController) Router{
	return &UserRouter{
		userController: _userController,
	}
} 

func (ur *UserRouter) Register(r chi.Router){
	r.Post("/signup", ur.userController.RegisterUser)
	r.With(middleware.AuthoriationMiddleware).Get("/{id}", ur.userController.FetchUserbyId)

	r.With(middleware.UserLoginRequestValidator).Post("/login", ur.userController.LoginUser)
	r.Post("/fetch_id", ur.userController.FetchUserbyId)
	r.Get("/getall_users", ur.userController.FetchAllUser)
}