package utils

import (
	"crypto/rand"
	"encoding/hex"
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
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}