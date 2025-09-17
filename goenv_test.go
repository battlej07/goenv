package goenv_test

import (
	"math"
	"testing"
	"time"

	"github.com/battlej07/goenv"
)

/* ---------- helpers ---------- */

func almostEq32(a, b float32, eps float32) bool { return float32(math.Abs(float64(a-b))) <= eps }
func almostEq64(a, b float64, eps float64) bool { return math.Abs(a-b) <= eps }

func expectPanic(t *testing.T, want bool) func() {
	t.Helper()
	return func() {
		if r := recover(); r != nil {
			if !want {
				t.Fatalf("unexpected panic: %v", r)
			}
		} else if want {
			t.Fatalf("expected panic, but none occurred")
		}
	}
}

/* ---------- string ---------- */

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback string
		want     string
	}{
		{name: "Env exists", key: "TEST_ENV_KEY", set: true, value: "value", fallback: "fallback", want: "value"},
		{name: "Env empty -> fallback", key: "TEST_ENV_KEY", set: true, value: "", fallback: "fallback", want: "fallback"},
		{name: "Env missing -> fallback", key: "TEST_ENV_KEY", set: false, fallback: "fallback", want: "fallback"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got := goenv.GetEnv(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnv(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    string
		wantErr bool
	}{
		{name: "Env exists", key: "TEST_ENV_KEY", set: true, value: "value", want: "value"},
		{name: "Env empty -> error", key: "TEST_ENV_KEY", set: true, value: "", wantErr: true},
		{name: "Env missing -> error", key: "TEST_ENV_KEY", set: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, gotErr := goenv.TryGetEnv(tt.key)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("TryGetEnv() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("TryGetEnv() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("TryGetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      string
		wantPanic bool
	}{
		{name: "Env exists", key: "TEST_ENV_KEY", set: true, value: "value", want: "value"},
		{name: "Env empty -> panic", key: "TEST_ENV_KEY", set: true, value: "", wantPanic: true},
		{name: "Env missing -> panic", key: "TEST_ENV_KEY", set: false, wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			got := goenv.MustGetEnv(tt.key)
			if got != tt.want {
				t.Errorf("MustGetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* ---------- int ---------- */

func TestGetEnvInt(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback int
		want     int
	}{
		{name: "Env exists int", key: "TEST_ENV_INT", set: true, value: "42", fallback: 7, want: 42},
		{name: "Env empty -> fallback", key: "TEST_ENV_INT", set: true, value: "", fallback: 7, want: 7},
		{name: "Env missing -> fallback", key: "TEST_ENV_INT", set: false, fallback: 7, want: 7},
		{name: "Env invalid -> fallback", key: "TEST_ENV_INT", set: true, value: "nope", fallback: 7, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			if got := goenv.GetEnvInt(tt.key, tt.fallback); got != tt.want {
				t.Errorf("GetEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvInt(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    int
		wantErr bool
	}{
		{name: "ok", key: "TRY_ENV_INT", set: true, value: "123", want: 123},
		{name: "empty -> err", key: "TRY_ENV_INT", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_ENV_INT", set: false, wantErr: true},
		{name: "bad int -> err", key: "TRY_ENV_INT", set: true, value: "x", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvInt(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("TryGetEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvInt(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      int
		wantPanic bool
	}{
		{name: "ok", key: "MUST_ENV_INT", set: true, value: "5", want: 5},
		{name: "empty -> panic", key: "MUST_ENV_INT", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_ENV_INT", set: false, wantPanic: true},
		{name: "bad int -> panic", key: "MUST_ENV_INT", set: true, value: "x", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				if got := goenv.MustGetEnvInt(tt.key); got != tt.want {
					t.Errorf("MustGetEnvInt() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvInt(tt.key)
			}
		})
	}
}

/* ---------- float32 ---------- */

func TestGetEnvFloat32(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback float32
		want     float32
	}{
		{name: "ok", key: "ENV_F32", set: true, value: "3.5", fallback: 1.1, want: 3.5},
		{name: "empty -> fallback", key: "ENV_F32", set: true, value: "", fallback: 1.1, want: 1.1},
		{name: "missing -> fallback", key: "ENV_F32", set: false, fallback: 1.1, want: 1.1},
		{name: "bad -> fallback", key: "ENV_F32", set: true, value: "x", fallback: 1.1, want: 1.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got := goenv.GetEnvFloat32(tt.key, tt.fallback)
			if !almostEq32(got, tt.want, 1e-6) {
				t.Errorf("GetEnvFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvFloat32(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    float32
		wantErr bool
	}{
		{name: "ok", key: "TRY_F32", set: true, value: "2.25", want: 2.25},
		{name: "empty -> err", key: "TRY_F32", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_F32", set: false, wantErr: true},
		{name: "bad -> err", key: "TRY_F32", set: true, value: "no", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvFloat32(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvFloat32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !almostEq32(got, tt.want, 1e-6) {
				t.Errorf("TryGetEnvFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvFloat32(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      float32
		wantPanic bool
	}{
		{name: "ok", key: "MUST_F32", set: true, value: "9.75", want: 9.75},
		{name: "empty -> panic", key: "MUST_F32", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_F32", set: false, wantPanic: true},
		{name: "bad -> panic", key: "MUST_F32", set: true, value: "bad", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				got := goenv.MustGetEnvFloat32(tt.key)
				if !almostEq32(got, tt.want, 1e-6) {
					t.Errorf("MustGetEnvFloat32() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvFloat32(tt.key)
			}
		})
	}
}

/* ---------- float64 ---------- */

func TestGetEnvFloat64(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback float64
		want     float64
	}{
		{name: "ok", key: "ENV_F64", set: true, value: "3.14159", fallback: 1.0, want: 3.14159},
		{name: "empty -> fallback", key: "ENV_F64", set: true, value: "", fallback: 1.0, want: 1.0},
		{name: "missing -> fallback", key: "ENV_F64", set: false, fallback: 1.0, want: 1.0},
		{name: "bad -> fallback", key: "ENV_F64", set: true, value: "bad", fallback: 1.0, want: 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got := goenv.GetEnvFloat64(tt.key, tt.fallback)
			if !almostEq64(got, tt.want, 1e-12) {
				t.Errorf("GetEnvFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvFloat64(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    float64
		wantErr bool
	}{
		{name: "ok", key: "TRY_F64", set: true, value: "6.28", want: 6.28},
		{name: "empty -> err", key: "TRY_F64", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_F64", set: false, wantErr: true},
		{name: "bad -> err", key: "TRY_F64", set: true, value: "no", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvFloat64(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !almostEq64(got, tt.want, 1e-12) {
				t.Errorf("TryGetEnvFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvFloat64(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      float64
		wantPanic bool
	}{
		{name: "ok", key: "MUST_F64", set: true, value: "1.25", want: 1.25},
		{name: "empty -> panic", key: "MUST_F64", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_F64", set: false, wantPanic: true},
		{name: "bad -> panic", key: "MUST_F64", set: true, value: "bad", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				got := goenv.MustGetEnvFloat64(tt.key)
				if !almostEq64(got, tt.want, 1e-12) {
					t.Errorf("MustGetEnvFloat64() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvFloat64(tt.key)
			}
		})
	}
}

/* ---------- bool ---------- */

func TestGetEnvBool(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback bool
		want     bool
	}{
		{name: "ok true", key: "ENV_BOOL", set: true, value: "true", fallback: false, want: true},
		{name: "empty -> fallback", key: "ENV_BOOL", set: true, value: "", fallback: true, want: true},
		{name: "missing -> fallback", key: "ENV_BOOL", set: false, fallback: true, want: true},
		{name: "bad -> fallback", key: "ENV_BOOL", set: true, value: "nah", fallback: false, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			if got := goenv.GetEnvBool(tt.key, tt.fallback); got != tt.want {
				t.Errorf("GetEnvBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvBool(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    bool
		wantErr bool
	}{
		{name: "ok true", key: "TRY_BOOL", set: true, value: "true", want: true},
		{name: "ok false", key: "TRY_BOOL", set: true, value: "false", want: false},
		{name: "empty -> err", key: "TRY_BOOL", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_BOOL", set: false, wantErr: true},
		{name: "bad -> err", key: "TRY_BOOL", set: true, value: "nope", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvBool(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvBool() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("TryGetEnvBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvBool(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      bool
		wantPanic bool
	}{
		{name: "ok", key: "MUST_BOOL", set: true, value: "true", want: true},
		{name: "empty -> panic", key: "MUST_BOOL", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_BOOL", set: false, wantPanic: true},
		{name: "bad -> panic", key: "MUST_BOOL", set: true, value: "??", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				if got := goenv.MustGetEnvBool(tt.key); got != tt.want {
					t.Errorf("MustGetEnvBool() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvBool(tt.key)
			}
		})
	}
}

/* ---------- time.Time (RFC3339) ---------- */

func TestGetEnvTime(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback time.Time
		want     time.Time
	}{
		{
			name:     "ok",
			key:      "ENV_TIME",
			set:      true,
			value:    "2025-08-24T12:34:56Z",
			fallback: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			want:     time.Date(2025, 8, 24, 12, 34, 56, 0, time.UTC),
		},
		{name: "empty -> fallback", key: "ENV_TIME", set: true, value: "", fallback: time.Unix(100, 0), want: time.Unix(100, 0)},
		{name: "missing -> fallback", key: "ENV_TIME", set: false, fallback: time.Unix(200, 0), want: time.Unix(200, 0)},
		{name: "bad -> fallback", key: "ENV_TIME", set: true, value: "notatime", fallback: time.Unix(300, 0), want: time.Unix(300, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got := goenv.GetEnvTime(tt.key, tt.fallback)
			if !got.Equal(tt.want) {
				t.Errorf("GetEnvTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvTime(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    time.Time
		wantErr bool
	}{
		{name: "ok", key: "TRY_TIME", set: true, value: "2025-08-24T00:00:00Z", want: time.Date(2025, 8, 24, 0, 0, 0, 0, time.UTC)},
		{name: "empty -> err", key: "TRY_TIME", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_TIME", set: false, wantErr: true},
		{name: "bad -> err", key: "TRY_TIME", set: true, value: "nope", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvTime(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !got.Equal(tt.want) {
				t.Errorf("TryGetEnvTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvTime(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      time.Time
		wantPanic bool
	}{
		{name: "ok", key: "MUST_TIME", set: true, value: "1999-12-31T23:59:59Z", want: time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC)},
		{name: "empty -> panic", key: "MUST_TIME", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_TIME", set: false, wantPanic: true},
		{name: "bad -> panic", key: "MUST_TIME", set: true, value: "bad", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				got := goenv.MustGetEnvTime(tt.key)
				if !got.Equal(tt.want) {
					t.Errorf("MustGetEnvTime() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvTime(tt.key)
			}
		})
	}
}

/* ---------- time.Duration ---------- */

func TestGetEnvDuration(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		set      bool
		value    string
		fallback time.Duration
		want     time.Duration
	}{
		{name: "ok", key: "ENV_DUR", set: true, value: "1h30m", fallback: 5 * time.Minute, want: 90 * time.Minute},
		{name: "empty -> fallback", key: "ENV_DUR", set: true, value: "", fallback: 5 * time.Minute, want: 5 * time.Minute},
		{name: "missing -> fallback", key: "ENV_DUR", set: false, fallback: 7 * time.Second, want: 7 * time.Second},
		{name: "bad -> fallback", key: "ENV_DUR", set: true, value: "nope", fallback: 10 * time.Second, want: 10 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			if got := goenv.GetEnvDuration(tt.key, tt.fallback); got != tt.want {
				t.Errorf("GetEnvDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGetEnvDuration(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		set     bool
		value   string
		want    time.Duration
		wantErr bool
	}{
		{name: "ok", key: "TRY_DUR", set: true, value: "250ms", want: 250 * time.Millisecond},
		{name: "empty -> err", key: "TRY_DUR", set: true, value: "", wantErr: true},
		{name: "missing -> err", key: "TRY_DUR", set: false, wantErr: true},
		{name: "bad -> err", key: "TRY_DUR", set: true, value: "zzz", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			got, err := goenv.TryGetEnvDuration(tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("TryGetEnvDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("TryGetEnvDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustGetEnvDuration(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		set       bool
		value     string
		want      time.Duration
		wantPanic bool
	}{
		{name: "ok", key: "MUST_DUR", set: true, value: "42s", want: 42 * time.Second},
		{name: "empty -> panic", key: "MUST_DUR", set: true, value: "", wantPanic: true},
		{name: "missing -> panic", key: "MUST_DUR", set: false, wantPanic: true},
		{name: "bad -> panic", key: "MUST_DUR", set: true, value: "oops", wantPanic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				t.Setenv(tt.key, tt.value)
			}
			defer expectPanic(t, tt.wantPanic)()
			if !tt.wantPanic {
				if got := goenv.MustGetEnvDuration(tt.key); got != tt.want {
					t.Errorf("MustGetEnvDuration() = %v, want %v", got, tt.want)
				}
			} else {
				_ = goenv.MustGetEnvDuration(tt.key)
			}
		})
	}
}
