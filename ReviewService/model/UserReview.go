package model

import "time"


type UserReview struct {
	Id         int64
	UserId     int64
	HotelId    int64
	BookingId  int64
	Comment    string
	Rating     float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}