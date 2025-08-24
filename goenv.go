// Package goenv provides helpers to load environment variables
// with support for type conversion, fallbacks, error returns, and
// panic-on-error variants.
package goenv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetEnv returns the value of the environment variable named by key.
// If the variable is unset or empty, it returns fallback.
func GetEnv(key, fallback string) string {
	v, err := TryGetEnv(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvInt returns the integer value of the environment variable named by key.
// If the variable is unset, empty, or cannot be parsed, it returns fallback.
func GetEnvInt(key string, fallback int) int {
	v, err := TryGetEnvInt(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvFloat32 returns the float32 value of the environment variable named by key.
// If the variable is unset, empty, or cannot be parsed, it returns fallback.
func GetEnvFloat32(key string, fallback float32) float32 {
	v, err := TryGetEnvFloat32(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvFloat64 returns the float64 value of the environment variable named by key.
// If the variable is unset, empty, or cannot be parsed, it returns fallback.
func GetEnvFloat64(key string, fallback float64) float64 {
	v, err := TryGetEnvFloat64(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvBool returns the boolean value of the environment variable named by key.
// If the variable is unset, empty, or cannot be parsed, it returns fallback.
func GetEnvBool(key string, fallback bool) bool {
	v, err := TryGetEnvBool(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvTime returns the time value of the environment variable named by key.
// The value must be in RFC3339 format. If the variable is unset, empty, or
// cannot be parsed, it returns fallback.
func GetEnvTime(key string, fallback time.Time) time.Time {
	v, err := TryGetEnvTime(key)
	if err != nil {
		return fallback
	}
	return v
}

// GetEnvDuration returns the duration value of the environment variable named by key.
// The value must be a valid time.ParseDuration string. If the variable is unset, empty,
// or cannot be parsed, it returns fallback.
func GetEnvDuration(key string, fallback time.Duration) time.Duration {
	v, err := TryGetEnvDuration(key)
	if err != nil {
		return fallback
	}
	return v
}

// TryGetEnv returns the value of the environment variable named by key.
// It returns an error if the variable is unset or empty.
func TryGetEnv(key string) (string, error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvInt returns the integer value of the environment variable named by key.
// It returns an error if the variable is unset, empty, or cannot be parsed as int.
func TryGetEnvInt(key string) (int, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("unable to convert %s to an integer", v)
		}
		return i, nil
	}
	return 0, fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvFloat32 returns the float32 value of the environment variable named by key.
// It returns an error if the variable is unset, empty, or cannot be parsed as float32.
func TryGetEnvFloat32(key string) (float32, error) {
	if v := os.Getenv(key); v != "" {
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0, fmt.Errorf("unable to convert %q to float32: %w", v, err)
		}
		return float32(f), nil
	}
	return 0, fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvFloat64 returns the float64 value of the environment variable named by key.
// It returns an error if the variable is unset, empty, or cannot be parsed as float64.
func TryGetEnvFloat64(key string) (float64, error) {
	if v := os.Getenv(key); v != "" {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to convert %q to float64: %w", v, err)
		}
		return f, nil
	}
	return 0, fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvBool returns the boolean value of the environment variable named by key.
// It returns an error if the variable is unset, empty, or cannot be parsed as bool.
func TryGetEnvBool(key string) (bool, error) {
	if v := os.Getenv(key); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf("unable to convert %q to bool: %w", v, err)
		}
		return b, nil
	}
	return false, fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvTime returns the time value of the environment variable named by key.
// The value must be in RFC3339 format. It returns an error if the variable is unset,
// empty, or cannot be parsed.
func TryGetEnvTime(key string) (time.Time, error) {
	if v := os.Getenv(key); v != "" {
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse %q as time (RFC3339): %w", v, err)
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unable to find env variable with key %s", key)
}

// TryGetEnvDuration returns the duration value of the environment variable named by key.
// The value must be a valid time.ParseDuration string. It returns an error if the variable
// is unset, empty, or cannot be parsed.
func TryGetEnvDuration(key string) (time.Duration, error) {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %q as duration: %w", v, err)
		}
		return d, nil
	}
	return 0, fmt.Errorf("unable to find env variable with key %s", key)
}

// MustGetEnv returns the value of the environment variable named by key.
// It panics if the variable is unset or empty.
func MustGetEnv(key string) string {
	v, err := TryGetEnv(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvInt returns the integer value of the environment variable named by key.
// It panics if the variable is unset, empty, or cannot be parsed as int.
func MustGetEnvInt(key string) int {
	v, err := TryGetEnvInt(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvFloat32 returns the float32 value of the environment variable named by key.
// It panics if the variable is unset, empty, or cannot be parsed as float32.
func MustGetEnvFloat32(key string) float32 {
	v, err := TryGetEnvFloat32(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvFloat64 returns the float64 value of the environment variable named by key.
// It panics if the variable is unset, empty, or cannot be parsed as float64.
func MustGetEnvFloat64(key string) float64 {
	v, err := TryGetEnvFloat64(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvBool returns the boolean value of the environment variable named by key.
// It panics if the variable is unset, empty, or cannot be parsed as bool.
func MustGetEnvBool(key string) bool {
	v, err := TryGetEnvBool(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvTime returns the time value of the environment variable named by key.
// The value must be in RFC3339 format. It panics if the variable is unset, empty,
// or cannot be parsed.
func MustGetEnvTime(key string) time.Time {
	v, err := TryGetEnvTime(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvDuration returns the duration value of the environment variable named by key.
// The value must be a valid time.ParseDuration string. It panics if the variable is unset,
// empty, or cannot be parsed.
func MustGetEnvDuration(key string) time.Duration {
	v, err := TryGetEnvDuration(key)
	if err != nil {
		panic(err)
	}
	return v
}
