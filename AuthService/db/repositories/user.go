package db

import "fmt"

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		// db : db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("This is repo layer")
	return nil
}