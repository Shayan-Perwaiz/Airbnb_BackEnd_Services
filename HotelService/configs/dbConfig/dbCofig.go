package config

import (
	config "GoHotelService/configs/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)


func SetUpDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = config.GetString("DBUSER", "root")
	cfg.Passwd = config.GetString("DBPASS", "")
	cfg.Net = "tcp"
	cfg.Addr = config.GetString("DBADDR", "localhost")
	cfg.DBName = config.GetString("DBNAME", "hotel_dev")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil{
		return nil, err
	}

	pingErr := db.Ping() //confirming if the db is connected 
	if pingErr != nil{
		fmt.Println("Error connecting the database")
		return nil, pingErr
	}
	fmt.Println("connected to the database")
	return db, nil


}