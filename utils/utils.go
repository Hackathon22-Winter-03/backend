package utils

import "os"

func GetEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v == "" {
		return defaultVal
	} else {
		return v
	}
}
