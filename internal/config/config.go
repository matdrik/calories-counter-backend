package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DBConnectionString string
	Host               string
	Port               string
}

func New() *Config {
	return &Config{
		DBConnectionString: getEnv("DB_CONNECTION_STRING", ""),
		Host:               getEnv("HOST", ""),
		Port:               getEnv("PORT", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("[config] Не найден параметр %v. Применено значение по-умолчанию.", key)
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
