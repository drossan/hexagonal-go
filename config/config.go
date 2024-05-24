package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Email    EmailConfig
}

type ServerConfig struct {
	Address   string
	JWTSecret string
}

type DatabaseConfig struct {
	URL string
}

type AppConfig struct {
	Env string
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
}

func LoadConfig() *Config {
	// Cargar variables de entorno desde el archivo .env si está en local
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found")
	}

	config := &Config{
		App: AppConfig{
			Env: os.Getenv("ENV"),
		},
		Server: ServerConfig{
			Address:   os.Getenv("SERVER_ADDRESS"),
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		Database: DatabaseConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		Email: EmailConfig{
			SMTPHost:     os.Getenv("SMTP_HOST"),
			SMTPPort:     os.Getenv("SMTP_PORT"),
			SMTPUser:     os.Getenv("SMTP_USER"),
			SMTPPassword: os.Getenv("SMTP_PASSWORD"),
			FromEmail:    os.Getenv("FROM_EMAIL"),
		},
	}

	return config
}
