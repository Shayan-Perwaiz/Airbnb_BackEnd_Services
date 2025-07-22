package db

import (
	models "GoAuth/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	Create(username string, email string, password string) error
	GetUserByEmail(email string)  (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db : db,
	}
}

func (u *UserRepositoryImpl) Create(username string, email string, password string) error {
	fmt.Println("This is repo layer")
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	user, err := u.db.Exec(query, username, email, password)
	if err != nil{
		fmt.Printf("Error in query")
		return fmt.Errorf("error inserting the user: %v", err)

	}
	rowsAffected, err := user.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("User is not created in storage")
		return fmt.Errorf("user is not created properly %v", err)
	}

	fmt.Println("User created")
	
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

func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error){
	query := "SELECT id, username, email , password, created_at FROM users WHERE email = ?"
	row := u.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_At)
	if err != nil{
		if err == sql.ErrNoRows{
		fmt.Println("No such user is found with this email:", email)
		return nil, fmt.Errorf("no user found with email: %s", email)
		}
		fmt.Println("Error fetching user:", err)
		return nil, fmt.Errorf("error fetching user: %w", err)

	}
	fmt.Println("User fetched successfully with email", email)
	return user, nil
}

func (u *UserRepositoryImpl) GetUserByID(id int64) (*models.User, error) {
    query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"
    row := u.db.QueryRow(query, id)

    user := &models.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_At, &user.Updated_At)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No user found with this ID:", id)
            return nil, fmt.Errorf("no user found with id: %d", id)
        }
        fmt.Println("Error fetching user:", err)
        return nil, fmt.Errorf("error fetching user: %w", err)
    }

    fmt.Println("User fetched successfully with ID:", id)
    return user, nil
}
