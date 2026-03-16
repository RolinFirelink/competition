package internal

import "os"

// GetEnv wraps os.Getenv for easier testing/mocking.
func GetEnv(key string) string {
	return os.Getenv(key)
}
