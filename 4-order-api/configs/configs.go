package configs

import (
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
		panic(".env file not found")
	}

	return &Config{
		DbConfig: DbConfig{Dsn: os.Getenv("DSN")},
		Address:  os.Getenv("ADDRESS"),
	}
}
