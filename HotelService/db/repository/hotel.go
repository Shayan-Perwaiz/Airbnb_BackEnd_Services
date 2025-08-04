package db

import (
	model "GoHotelService/models"
	"database/sql"
	"fmt"
)

type HotelRepository interface {
	FindHotelById(hotelId int64) (*model.Hotel, error)
	HotelByName(hotelName string) (*model.Hotel, error)
	DeleteById(id int64) error
}

type HotelRepositoryImpl struct {
	db *sql.DB
}

func NewHotelRepositoryImpl(_db *sql.DB) HotelRepository{
	return &HotelRepositoryImpl{
		db : _db,
	}
}

func(repo *HotelRepositoryImpl) FindHotelById(hotelId int64) (*model.Hotel, error){
	query := "SELECT id FROM hotels WHERE id = ?"
	row := repo.db.QueryRow(query, hotelId)
	hotel := &model.Hotel{}
	err := row.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.City, &hotel.CreatedAt, &hotel.UpdatedAt)
	if err == sql.ErrNoRows{
		fmt.Println("no such row found with id", hotelId)
		return nil, err
	}
	fmt.Println("hotel by id", hotelId)
	return hotel, nil

}

func(repo *HotelRepositoryImpl) HotelByName(hotelName string) (*model.Hotel, error){
	query := "SELECT name FROM hotels WHERE name = ?"
	row := repo.db.QueryRow(query, hotelName)
	hotel := &model.Hotel{}
	err := row.Scan(&hotel.Id, &hotel.Name, &hotel.City, &hotel.CreatedAt, &hotel.UpdatedAt)
	if err != sql.ErrNoRows{
		fmt.Println("no hotel found with name", hotelName)
		return nil, err
	}
	fmt.Println("hotel by name", hotelName)
	return hotel, nil
}

func(repo *HotelRepositoryImpl) GetAll(){

}

func(repo *HotelRepositoryImpl) DeleteById(id int64) error{
	query := "DELETE FROM hotels WHERE ID = ?"
	result, err := repo.db.Exec(query, id)
	if err != nil{
		fmt.Println("falied to execute sql delete query")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil{
			fmt.Println("failed to retrieve affected rows")
		}
	if rowsAffected == 0{
		fmt.Println("no hotel found with id", id)
	}else{
		fmt.Printf("successfully deleted hotel with ID %d\n", id)
	}
	return nil

}

