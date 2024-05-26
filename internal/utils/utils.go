package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
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

func HashPass(pw string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Errorf("can't hash password: %w", err))
	}
	return string(hashed)
}

func CheckPass(raw string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(raw), []byte(hashed)) != nil
}
