package config

import (
	"os"
)

type Config struct {
	PostgresDSN string
	RedisAddr   string
	RedisPass   string
	RedisDB     int
	HTTPPort    string
}

func LoadConfig() *Config {
	return &Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RedisPass:   os.Getenv("REDIS_PASS"),
		RedisDB:     0, // можно добавить чтение из env
		HTTPPort:    os.Getenv("HTTP_PORT"),
	}
}

/*
func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
*/
