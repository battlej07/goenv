// Package goenv provides helpers to load environment variables
// with support for type conversion, fallbacks, error returns, and
// panic-on-error variants.
package goenv

import (
	"fmt"
	"os"
	"reflect"
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

// Load populates a struct's fields from environment variables using struct tags.
// Each field should have a `goenv:"ENV_VAR_NAME"` tag to specify which environment
// variable to load. Optionally, a `fallback:"value"` tag can be used to provide a
// default value if the environment variable is missing or invalid.
// The function uses reflection to set field values based on their types.
// The input must be a pointer to a struct. Returns an error if the input is invalid
// or if any required environment variable cannot be loaded (and no fallback is provided).
func Load(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return fmt.Errorf("Load expects a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Load expects a pointer to a struct, got %s", val.Kind())
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if !field.CanSet() {
			continue
		}

		tag := fieldType.Tag.Get("goenv")
		if tag == "" {
			continue
		}

		fallbackTag := fieldType.Tag.Get("fallback")
		if err := setField(field, tag, fallbackTag); err != nil {
			return fmt.Errorf("field %s: %w", fieldType.Name, err)
		}
	}

	return nil
}

func setField(field reflect.Value, envKey, fallback string) error {
	switch field.Kind() {
	case reflect.String:
		v, err := TryGetEnv(envKey)
		if err != nil {
			if fallback != "" {
				field.SetString(fallback)
				return nil
			}
			return err
		}
		field.SetString(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Type() == reflect.TypeFor[time.Duration]() {
			v, err := TryGetEnvDuration(envKey)
			if err != nil {
				if fallback != "" {
					d, parseErr := time.ParseDuration(fallback)
					if parseErr != nil {
						return fmt.Errorf("invalid fallback duration %q: %w", fallback, parseErr)
					}
					field.SetInt(int64(d))
					return nil
				}
				return err
			}
			field.SetInt(int64(v))
		} else {
			v, err := TryGetEnvInt(envKey)
			if err != nil {
				if fallback != "" {
					i, parseErr := strconv.Atoi(fallback)
					if parseErr != nil {
						return fmt.Errorf("invalid fallback integer %q: %w", fallback, parseErr)
					}
					field.SetInt(int64(i))
					return nil
				}
				return err
			}
			field.SetInt(int64(v))
		}

	case reflect.Float32:
		v, err := TryGetEnvFloat32(envKey)
		if err != nil {
			if fallback != "" {
				f, parseErr := strconv.ParseFloat(fallback, 32)
				if parseErr != nil {
					return fmt.Errorf("invalid fallback float32 %q: %w", fallback, parseErr)
				}
				field.SetFloat(f)
				return nil
			}
			return err
		}
		field.SetFloat(float64(v))

	case reflect.Float64:
		v, err := TryGetEnvFloat64(envKey)
		if err != nil {
			if fallback != "" {
				f, parseErr := strconv.ParseFloat(fallback, 64)
				if parseErr != nil {
					return fmt.Errorf("invalid fallback float64 %q: %w", fallback, parseErr)
				}
				field.SetFloat(f)
				return nil
			}
			return err
		}
		field.SetFloat(v)

	case reflect.Bool:
		v, err := TryGetEnvBool(envKey)
		if err != nil {
			if fallback != "" {
				b, parseErr := strconv.ParseBool(fallback)
				if parseErr != nil {
					return fmt.Errorf("invalid fallback bool %q: %w", fallback, parseErr)
				}
				field.SetBool(b)
				return nil
			}
			return err
		}
		field.SetBool(v)

	case reflect.Struct:
		if field.Type() == reflect.TypeFor[time.Time]() {
			v, err := TryGetEnvTime(envKey)
			if err != nil {
				if fallback != "" {
					t, parseErr := time.Parse(time.RFC3339, fallback)
					if parseErr != nil {
						return fmt.Errorf("invalid fallback time %q (RFC3339): %w", fallback, parseErr)
					}
					field.Set(reflect.ValueOf(t))
					return nil
				}
				return err
			}
			field.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("unsupported struct type %s", field.Type())
		}

	default:
		return fmt.Errorf("unsupported field type %s", field.Kind())
	}

	return nil
}
