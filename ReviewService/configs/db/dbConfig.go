package config

import (
	config "GoReview/configs/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetUpDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = config.GetString("DB_USER", "root")
	cfg.Passwd = config.GetString("DB_PASSWORD", "db_password")
	cfg.Net = "tcp"
	cfg.Addr = config.GetString("DB_ADDR", "127.0.0.1:3306")
	cfg.DBName = config.GetString("DB_NAME", "review_dev")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil{
		fmt.Println("Error, database configuration values")
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil{
		fmt.Println("fails to connect databse")
		return nil, err
	}
	fmt.Println("Connected Successfully to the database")
	return db, nil
}