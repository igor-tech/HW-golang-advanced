package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type Config struct {
	DbConfig
	Address string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	return &Config{
		DbConfig: DbConfig{Dsn: os.Getenv("DSN")},
		Address:  os.Getenv("ADDRESS"),
	}
}
