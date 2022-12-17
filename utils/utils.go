package utils

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

var (
	ErrNotFound error = fmt.Errorf("not found")
)

func GetEnv(key, defaultVal string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	return v
}

// generateID uniqueなIDを生成する
func GenerateID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
