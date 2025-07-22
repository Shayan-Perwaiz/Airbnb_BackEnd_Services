package services

import (
	configEnv "GoAuth/Config/env"
	db "GoAuth/db/repositories"
	util "GoAuth/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)


type UserService interface {
	CreateUser() error
	LoginUser() (string,error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService{ 
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (us *UserServiceImpl) CreateUser() error{
	fmt.Println("Creating user in User Service")
	hashedPassword,_ := util.HashedPassword("filziperwaiz")

	us.userRepository.Create("filzi", "filzi@123gmail.com", hashedPassword)
	return nil
}

func (us *UserServiceImpl) LoginUser() (string,error){
	email := "filzi@123gmail.com"
	textpassword := "filziperwaiz"
	userDetail, err := us.userRepository.GetUserByEmail(email)
	if err != nil{
		fmt.Println("Error fetching the suer")
		return "", err
	}
	if userDetail == nil {
		fmt.Print("No user found with the given email")
		return "",err
	}
	isValidPassword := util.ValidatePassword(userDetail.Password, textpassword)

	if !isValidPassword{
		fmt.Println("Password does not match")
	}

	fmt.Println("Here is the user details below: ")
	fmt.Println(userDetail.ID)
	fmt.Println(userDetail.Username)
	fmt.Println(userDetail.Email)
	fmt.Println(userDetail.Password)
	fmt.Println(userDetail.Created_At)

	paylaod := jwt.MapClaims{
		"user_id" : userDetail.ID,
		"user_email" : userDetail.Email,
		"user_name": userDetail.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, paylaod)
	tokenString, err := token.SignedString([]byte(configEnv.GetString("JWT_TOKEN", "TOKEN")))
	if err != nil{
		fmt.Println("error creating a JWT token")
		return  "",err
	}
	fmt.Println("JWT Token :",tokenString)

	return tokenString, nil
}

