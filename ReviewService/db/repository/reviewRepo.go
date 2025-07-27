package db

import (
	"GoReview/dto"
	"database/sql"
	"fmt"
	"time"
)

type ReviewRepository interface {
	CreateUserReview(payload *dto.UserReviewDto) error
}

type ReviewRepositoryImpl struct {
	db *sql.DB
}

func NewReviewRepositoryImpl(_db *sql.DB) *ReviewRepositoryImpl{
	return &ReviewRepositoryImpl{
		db : _db,
	}
}

func(rr *ReviewRepositoryImpl) CreateUserReview(payload *dto.UserReviewDto) error{
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
	// result, err := rr.db.Exec(query, user_id, hotel_id, booking_id, comment, rating, created_at, updated_at)
	result, err := rr.db.Exec(query, payload.UserID, payload.HotelID, payload.BookingId, payload.Comment, payload.Rating, time.Now(), time.Now())
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