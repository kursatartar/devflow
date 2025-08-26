package config

import (
	"os"
	"time"
)

type Config struct {
	Database DatabaseConfig
}
type DatabaseConfig struct {
	MongoURI string
	DBName   string
	MaxPool  uint64
	Timeout  time.Duration
}

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			MongoURI: getEnv("MONGO_URI", "mongodb://127.0.0.1:27017"),
		}}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
