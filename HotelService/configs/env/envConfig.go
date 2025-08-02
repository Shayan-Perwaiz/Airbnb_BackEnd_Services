package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil{
		fmt.Println("error loading the env variables")
	}
}

func GetString(key string, fallback string) string{
	value, ok := os.LookupEnv(key)
	if !ok{
		return fallback
	}
	return value
}

