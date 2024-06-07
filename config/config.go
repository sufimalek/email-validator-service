package config

import (
	"os"
)

var (
	Port string
)

func LoadConfig() {
	Port = getEnv("PORT", "8080")
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
