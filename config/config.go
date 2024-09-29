package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_PORT       string
	SERVER_API_PREFIX string

	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_PASSWORD string
	PG_DB_NAME  string
	PG_SCHEMA   string
	PG_SSL_MODE string

	HMAC_KEY string
}

var App = initConfig()

func initConfig() Config {
	godotenv.Load()

	config := Config{
		SERVER_PORT:       getEnv("SERVER_PORT", "8000"),
		SERVER_API_PREFIX: getEnv("SERVER_API_PREFIX", "/app"),

		PG_HOST:     getEnv("PG_HOST", "127.0.0.1"),
		PG_PORT:     getEnv("PG_PORT", "5432"),
		PG_USER:     getEnv("PG_USER", "app"),
		PG_PASSWORD: getEnv("PG_PASSWORD", "postgres-password"),
		PG_DB_NAME:  getEnv("PG_DB_NAME", "my-app"),
		PG_SCHEMA:   getEnv("PG_SCHEMA", "my-app"),
		PG_SSL_MODE: getEnv("PG_SSL_MODE", "disable"),
		HMAC_KEY:    getEnv("HMAC_KEY", "HMAC_KEY"),
	}

	log.Printf("debugging server config: %+v", config)
	log.Print("WARNING: should remove server config debugging logs in production")

	return config
}

func getEnv(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return def
}
