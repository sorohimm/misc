// Package env provides a utility functions to work with env
package env

import "os"

// Getenv returns the value of the environment variable key.
// If the variable is not set, it returns the default value, if provided.
// If the variable is not set and there is no default value, it returns an empty string.
func Getenv(key string, defaultValue ...string) string {
	env := os.Getenv(key)
	if env == "" {
		if defaultValue != nil {
			return defaultValue[0]
		}
		return ""
	}

	return env
}
