package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

func LoadEnvOrDefault(envName string, defaultVal string) string {
	val, found := os.LookupEnv(envName)
	if !found {
		return defaultVal
	}
	return val

}

func GenerateToken() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(fmt.Errorf("can't generate random token: %w", err))
	}
	return hex.EncodeToString(bytes)
}
