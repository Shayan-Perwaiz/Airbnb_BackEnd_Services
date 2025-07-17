package config

import (
	config "GoAuth/Config/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func SetUpDB() (*sql.DB, error){
	cfg := mysql.NewConfig()

	cfg.User = config.GetString("DB_USER", "user")
	cfg.Passwd = config.GetString("DB_PASSWORD", "pswd")
	cfg.Net = "tcp"
	cfg.Addr = config.GetString("DB_ADDR", "addr")
	cfg.DBName = config.GetString("DB_NAME", "dbname")
	
	fmt.Println("Connecting to database", cfg.DBName)

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil{
		fmt.Println("Error connecting the databse")
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
		return nil, pingErr 

	}
	fmt.Println("Connected to", cfg.DBName)
	return db, nil


}