package tools

import "os"

func GetEnv(name, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return defaultValue
}
