package db

import (
	models "GoAuth/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	Create() error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db : db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("This is repo layer")
	
	return nil
}

func(u *UserRepositoryImpl) GetAll() ([]*models.User, error){
	var usersAll []*models.User
	query := "SELECT * FROM users"
	rows, err := u.db.Query(query)
	if err != nil{
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next(){
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_At, &user.Updated_At)
		if err != nil{
			    return nil, err
			}
		usersAll = append(usersAll, user)
	}
	if err := rows.Err(); err != nil{
			return nil, fmt.Errorf("row iteration error: %v", err)
		}

	return usersAll, nil
}

func(u *UserRepositoryImpl) DeleteById(id int64) error{
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)
	if err != nil{
		return fmt.Errorf("failed to delete user: %v", err)
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil{
		return fmt.Errorf("could not check rows affected: %v", err)
	}
	if rowsAffected == 0{
		return fmt.Errorf("no user found with id %d", id)
	}
	fmt.Println("User successfully deleted with id", id)

	return nil
}
