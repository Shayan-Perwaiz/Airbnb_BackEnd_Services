package services

import (
	db "GoAuth/db/repositories"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


type UserService interface {
	CreateUser() error
	EncryptPassword(password string) error
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
	us.userRepository.Create()
	return nil
}

func(us *UserServiceImpl) EncryptPassword(password string) error{
	// plainText := []byte(password)
	sha256Hash := sha256.Sum256([]byte(password)) 
	fmt.Println(hex.EncodeToString(sha256Hash[:]))
	return nil
}