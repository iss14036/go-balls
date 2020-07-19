package config

import "os"

func GetString(key string) string {
	return os.Getenv(key)
}