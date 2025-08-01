package services

import (
	configEnv "GoAuth/Config/env"
	db "GoAuth/db/repositories"
	"GoAuth/dto"
	util "GoAuth/utils"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)


type UserService interface {
	CreateUser() error
	RemoveUserById(id int64) error
	LoginUser(payload *dto.LoginUserRequestDto) (string,error)
	FindUserById(id int64) error
	GetAllUsers() error
	
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
func (us *UserServiceImpl) RemoveUserById(id int64) error{
	return nil
}

func (us *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDto) (string,error){
	userDetail, err := us.userRepository.GetUserByEmail(payload.Email)
	if err != nil{
		fmt.Println("Error fetching the user")
		return "", errors.New("error fetching the user")
	}
	if userDetail == nil {
		return "",fmt.Errorf("no users found with this email: %s", payload.Email)
	}
	isValidPassword := util.ValidatePassword(userDetail.Password, payload.Password)

	if !isValidPassword{
		fmt.Println("Password does not match")
		return "", errors.New("invalid Password")
	}

	fmt.Println("Here is the user details below: ")
	// fmt.Println(userDetail.ID)
	// fmt.Println(userDetail.Username)
	// fmt.Println(userDetail.Email)
	// fmt.Println(userDetail.Password)
	// fmt.Println(userDetail.Created_At)
	fmt.Println(&userDetail)

	jwtPaylaod := jwt.MapClaims{
		"user_id" : userDetail.ID,
		"user_email" : userDetail.Email,
		"user_name": userDetail.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPaylaod)
	tokenString, err := token.SignedString([]byte(configEnv.GetString("JWT_TOKEN", "TOKEN")))
	if err != nil{
		fmt.Println("error creating a JWT token")
		return  "",errors.New("error creating a JWT token")
	}
	fmt.Println("JWT Token :",tokenString)

	return tokenString, nil
}

func(us *UserServiceImpl) FindUserById(id int64) error{
	user, err := us.userRepository.GetUserByID(id)
	if err != nil{
		fmt.Println("error fetching the user")
		return errors.New("error fetching the user")
	}

	if user == nil{
		fmt.Println("no user found with id", id)
		return fmt.Errorf("no users found with this id: %d", id)
	}
	fmt.Println(user)
	return nil
}

func(us *UserServiceImpl) GetAllUsers() error{
	users, err := us.userRepository.GetAll()
	if err != nil{
		fmt.Println("error fetching the users")
		return err
	}

	if users == nil{
		fmt.Println("no users found")
		return err
	}

	for _, user := range users {
    fmt.Println(user)
    }

	return nil
}

