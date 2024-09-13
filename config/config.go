package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST      string
	DB_PORT      int
	DB_PASSWORD  string
	DB_NAME      string
	DB_USER      string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "alert"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "salom"))
	config.DB_PORT = cast.ToInt(Coalesce("DB_PORT", 5432))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}