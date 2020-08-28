package utils

import (
	"os"
	"testing"
)

func TestGetEnvFallback(t *testing.T) {
	fallback := "test_Fallback_GetEnv"
	result := GetEnv("NON_EXISTENT_ENV_VARIABLE", fallback)
	if result != fallback {
		t.Errorf("Expected %s got %s", fallback, result)
	}
}

func TestGetEnv(t *testing.T) {
	value := "test_GetEnv"
	os.Setenv("TEST_ENV_VAR", value)
	result := GetEnv("TEST_ENV_VAR", "fallback_value")
	if result != value {
		t.Errorf("Expected %s got %s", value, result)
	}
}
