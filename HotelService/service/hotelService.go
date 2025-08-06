package services

import db "GoHotelService/db/repository"

type HotelService interface {
	GetAllHotels()
	AddHotel()
	GetHotelByid(id int64)
	GetHotelName(hotelName string)
	DeleteHotelId(id int64)
}

type HotelServiceImpl struct {
	hotelRepository db.HotelRepositoryImpl
}

func NewHotelServiceImpl(_hotelRepository db.HotelRepositoryImpl) HotelService{
	return &HotelServiceImpl{
		hotelRepository: _hotelRepository,
	}
}

func (service *HotelServiceImpl) GetAllHotels(){
	service.hotelRepository.GetAll()
}

func (service *HotelServiceImpl) AddHotel(){
	service.hotelRepository.InsertHotel()
}

func (service *HotelServiceImpl) GetHotelByid(id int64){
	service.hotelRepository.FindHotelById(id)
}

func (service *HotelServiceImpl) GetHotelName(hotelName string){
	service.hotelRepository.HotelByName(hotelName)
}

func (service *HotelServiceImpl) DeleteHotelId(id int64){
	service.hotelRepository.DeleteById(id)
}