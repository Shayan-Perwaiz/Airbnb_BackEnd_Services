package dto

type UserReviewDto struct {
	UserID    int64   `json:"user_id"`
	HotelID   int64   `json:"hotel_id"`
	BookingId int64   `json:"booking_id"`
	Comment   string  `json:"comment"`
	Rating    float64 `json:"rating"`
}