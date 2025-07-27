package db

import (
	"database/sql"
	"fmt"
	"time"
)

type ReviewRepository interface {
	CreateUserReview(user_id int64, hotel_id int64,
		booking_id int64, comment string,
		rating float64, created_at time.Time,
		updated_at time.Time) error
}

type ReviewRepositoryImpl struct {
	db *sql.DB
}

func NewReviewRepositoryImpl(_db *sql.DB) *ReviewRepositoryImpl{
	return &ReviewRepositoryImpl{
		db : _db,
	}
}

func(rr *ReviewRepositoryImpl) CreateUserReview(user_id int64, hotel_id int64,
	 booking_id int64, comment string, 
	 rating float64, created_at time.Time, 
	 updated_at time.Time) error{
	query := `
	INSERT INTO user_reviews (
		user_id,
		hotel_id,
		booking_id,
		comment,
		rating,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?)
`
	result, err := rr.db.Exec(query, user_id, hotel_id, booking_id, comment, rating, created_at, updated_at)
	if err != nil{
		fmt.Println("Error inserting values into user review table")
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil{
		fmt.Println("Error getting rows affected:", err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("no rows inserted")
	}

	fmt.Println("User review inserted successfully")
	return nil

}