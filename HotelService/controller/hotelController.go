package controllers

import (
	dto "GoHotelService/dtos"
	services "GoHotelService/service"
	"fmt"
	"net/http"
)

type HotelControllerImpl struct {
	hotelService services.HotelServiceImpl
}

func NewHotelControllerImpl(_hotelService services.HotelServiceImpl) *HotelControllerImpl{
	return &HotelControllerImpl{
		hotelService: _hotelService,
	}
}

func(controller *HotelControllerImpl) AddHotel(w http.ResponseWriter, r *http.Request){
	payload := dto.HotelDto{}

	fmt.Println("This is AddHotelController")

}