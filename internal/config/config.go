package config

import (
	"os"
	"strconv"
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
	cfg := &Config{
		Database: DatabaseConfig{
			MongoURI: getEnv("MONGO_URI", "mongodb://127.0.0.1:27017"),
			DBName:   getEnv("DB_NAME", ""),
			MaxPool:  parseUint64(os.Getenv("MONGO_MAX_POOL")),
			Timeout:  parseDuration(os.Getenv("MONGO_TIMEOUT")),
		},
	}

	if cfg.Database.DBName == "" {
		cfg.Database.DBName = "devflow"
	}
	if cfg.Database.MaxPool == 0 {
		cfg.Database.MaxPool = 10
	}
	if cfg.Database.Timeout == 0 {
		cfg.Database.Timeout = 10 * time.Second
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func parseUint64(s string) uint64 {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func parseDuration(s string) time.Duration {
	if s == "" {
		return 0
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0
	}
	return d
}
