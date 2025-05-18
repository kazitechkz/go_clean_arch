package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// GetString retrieves a string value from the environment
func GetString(key string) string {
	return os.Getenv(key)
}

// GetInt retrieves and parses an integer value from the environment
func GetInt(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("Неверное Цифровое значение для %s: %v", key, err)
	}
	return val
}

// GetBool retrieves and parses a boolean value from the environment
func GetBool(key string) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		log.Fatalf("Неверное булево значение %s: %v", key, err)
	}
	return val
}

// GetStringPtr retrieves a string pointer from the environment
func GetStringPtr(key string) *string {
	val := os.Getenv(key)
	return &val
}

// GetIntPtr retrieves an int pointer from the environment
func GetIntPtr(key string) *int {
	val := GetInt(key)
	return &val
}

func EscapeLikePattern(s string) string {
	s = strings.ReplaceAll(s, "_", "\\_")
	s = strings.ReplaceAll(s, "%", "\\%")
	return s
}
