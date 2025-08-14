package util

import "os"

// GetEnv
//
// Get env value by key.
func GetEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
