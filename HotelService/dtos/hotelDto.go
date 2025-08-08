package dto

type HotelDto struct {
	HotelName    string `json:"hotelname" validate:"required, min=3, max=20"`
	HotelAddress string `json:"hoteladdress" validate:"required, min=3, max=20"`
	HotelCity    string `json:"hotelcity" validate:"required, min=3, max=20"`
}