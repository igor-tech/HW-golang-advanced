package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Email    string
	Password string
	Address  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	return &Config{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
