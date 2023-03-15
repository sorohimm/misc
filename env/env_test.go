package env

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	key := "TEST_KEY"
	defaultValue := "default_value"

	// Test with an existing key
	os.Setenv(key, "existing_value")
	value := Getenv(key, defaultValue)
	if value != "existing_value" {
		t.Errorf("Expected 'existing_value', got '%s'", value)
	}

	// Test with a non-existing key and default value
	value = Getenv("NON_EXISTING_KEY", defaultValue)
	if value != defaultValue {
		t.Errorf("Expected '%s', got '%s'", defaultValue, value)
	}

	// Test with a non-existing key and no default value
	value = Getenv("NON_EXISTING_KEY")
	if value != "" {
		t.Errorf("Expected empty string, got '%s'", value)
	}
}
