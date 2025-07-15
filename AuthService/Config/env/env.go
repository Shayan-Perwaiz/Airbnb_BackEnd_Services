package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)
func Load(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println("Error loading the .env file")
	}
}

func GetString(key string, fallback string) string{
	value, ok := os.LookupEnv(key)
	if !ok{
		return fallback
	}
	return value
}