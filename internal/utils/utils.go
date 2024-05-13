package utils

import "os"

func LoadEnvOrDefault(envName string, defaultVal string) string {
	val, found := os.LookupEnv(envName)
	if !found {
		return defaultVal
	}
	return val
}
