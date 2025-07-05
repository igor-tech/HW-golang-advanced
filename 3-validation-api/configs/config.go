package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SMTPHost string
	SMTPPassword string
	SMTPEmail string
	SMTPPort string
	BaseUrl string
	Address string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	return &Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		SMTPEmail: os.Getenv("SMTP_EMAIL"),
		SMTPPort: os.Getenv("SMTP_PORT"),
		BaseUrl: os.Getenv("BASE_URL"),
		Address: os.Getenv("ADDRESS"),
	}
}
