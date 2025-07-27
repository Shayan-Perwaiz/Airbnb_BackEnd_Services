package model

import "time"


type UserReview struct {
	id         int64
	userId     int64
	hotelId    int64
	bookingId  int64
	comment    string
	rating     float64
	createdAt  time.Time
	updatedAt  time.Time
}