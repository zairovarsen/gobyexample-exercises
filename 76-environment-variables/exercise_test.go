package environmentvariables

import (
	"maps"
	"os"
	"strings"
	"testing"
)

// All variables used here start with GBE_ and are written in upper case
// only: on Windows the names of environment variables are case-insensitive,
// so names that differ only in case would be the same variable there.
//
// t.Setenv sets a variable for the duration of one test and restores the old
// state afterwards.

// unsetenv makes sure that key does not exist during the test.
func unsetenv(t *testing.T, key string) {
	t.Helper()
	t.Setenv(key, "") // registers the restore of the original state
	if err := os.Unsetenv(key); err != nil {
		t.Fatal(err)
	}
}

func TestGetOrDefault(t *testing.T) {
	t.Setenv("GBE_GET_SET", "1")
	t.Setenv("GBE_GET_EMPTY", "")
	unsetenv(t, "GBE_GET_UNSET")

	tests := []struct {
		key, def string
		want     string
	}{
		{"GBE_GET_SET", "fallback", "1"},
		{"GBE_GET_SET", "", "1"},
		{"GBE_GET_EMPTY", "fallback", "fallback"},
		{"GBE_GET_UNSET", "fallback", "fallback"},
		{"GBE_GET_UNSET", "", ""},
	}
	for _, tt := range tests {
		if got := GetOrDefault(tt.key, tt.def); got != tt.want {
			t.Errorf("GetOrDefault(%q, %q) = %q, want %q", tt.key, tt.def, got, tt.want)
		}
	}
}

func TestStatus(t *testing.T) {
	t.Setenv("GBE_STATUS_SET", "1")
	t.Setenv("GBE_STATUS_PAIR", "a=b")
	t.Setenv("GBE_STATUS_EMPTY", "")
	unsetenv(t, "GBE_STATUS_UNSET")

	tests := []struct {
		key  string
		want string
	}{
		{"GBE_STATUS_SET", "GBE_STATUS_SET=1"},
		{"GBE_STATUS_PAIR", "GBE_STATUS_PAIR=a=b"},
		{"GBE_STATUS_EMPTY", "GBE_STATUS_EMPTY is empty"},
		{"GBE_STATUS_UNSET", "GBE_STATUS_UNSET is not set"},
	}
	for _, tt := range tests {
		if got := Status(tt.key); got != tt.want {
			t.Errorf("Status(%q) = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestSetAll(t *testing.T) {
	vars := map[string]string{
		"GBE_SETALL_FOO":   "1",
		"GBE_SETALL_WORDS": "two words",
		"GBE_SETALL_PAIR":  "x=y",
	}
	for key := range vars {
		unsetenv(t, key)
	}

	if err := SetAll(vars); err != nil {
		t.Fatalf("SetAll(%v) returned error: %v", vars, err)
	}
	for key, want := range vars {
		got, ok := os.LookupEnv(key)
		if !ok {
			t.Errorf("after SetAll the variable %s is not set", key)
		} else if got != want {
			t.Errorf("after SetAll the variable %s is %q, want %q", key, got, want)
		}
	}

	if err := SetAll(nil); err != nil {
		t.Errorf("SetAll(nil) returned error: %v", err)
	}
	if err := SetAll(map[string]string{"": "no name"}); err == nil {
		t.Error(`SetAll({"": "no name"}) returned a nil error, want the error from os.Setenv for the empty key`)
	}
}

func TestEnvMap(t *testing.T) {
	t.Setenv("GBE_MAP_HOST", "localhost")
	t.Setenv("GBE_MAP_PORT", "80")
	t.Setenv("GBE_MAP_DSN", "postgres://db/app?sslmode=disable&timeout=5")
	t.Setenv("GBE_OTHER_NAME", "not included")

	want := map[string]string{
		"GBE_MAP_HOST": "localhost",
		"GBE_MAP_PORT": "80",
		"GBE_MAP_DSN":  "postgres://db/app?sslmode=disable&timeout=5",
	}
	if got := EnvMap("GBE_MAP_"); !maps.Equal(got, want) {
		t.Errorf("EnvMap(\"GBE_MAP_\") = %v, want %v", got, want)
	}

	if got := EnvMap("GBE_NO_SUCH_PREFIX_"); len(got) != 0 {
		t.Errorf("EnvMap(\"GBE_NO_SUCH_PREFIX_\") = %v, want an empty map", got)
	}

	// The whole environment: every variable set above must be in there.
	all := EnvMap("")
	for _, key := range []string{"GBE_MAP_HOST", "GBE_OTHER_NAME"} {
		if _, ok := all[key]; !ok {
			t.Errorf("EnvMap(\"\") does not contain %s, want all environment variables", key)
		}
	}
}

func TestLoadConfig(t *testing.T) {
	const prefix = "GBE_CFG_"
	names := []string{prefix + "HOST", prefix + "PORT", prefix + "DEBUG"}

	tests := []struct {
		env  map[string]string
		want Config
	}{
		{nil, Config{Host: "localhost", Port: 8080, Debug: false}},
		{
			map[string]string{"GBE_CFG_HOST": "example.com", "GBE_CFG_PORT": "9000", "GBE_CFG_DEBUG": "true"},
			Config{Host: "example.com", Port: 9000, Debug: true},
		},
		{
			map[string]string{"GBE_CFG_PORT": "443"},
			Config{Host: "localhost", Port: 443, Debug: false},
		},
		{
			map[string]string{"GBE_CFG_DEBUG": "1"},
			Config{Host: "localhost", Port: 8080, Debug: true},
		},
		{
			map[string]string{"GBE_CFG_HOST": "", "GBE_CFG_PORT": "", "GBE_CFG_DEBUG": "false"},
			Config{Host: "localhost", Port: 8080, Debug: false},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			for _, name := range names {
				unsetenv(t, name)
			}
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			got, err := LoadConfig(prefix)
			if err != nil {
				t.Fatalf("LoadConfig(%q) with env %v returned error: %v", prefix, tt.env, err)
			}
			if got != tt.want {
				t.Errorf("LoadConfig(%q) with env %v = %+v, want %+v", prefix, tt.env, got, tt.want)
			}
		})
	}

	invalid := []struct {
		key, value string
	}{
		{"GBE_CFG_PORT", "eighty"},
		{"GBE_CFG_DEBUG", "maybe"},
	}
	for _, tt := range invalid {
		t.Run("invalid "+tt.key, func(t *testing.T) {
			for _, name := range names {
				unsetenv(t, name)
			}
			t.Setenv(tt.key, tt.value)
			got, err := LoadConfig(prefix)
			if err == nil {
				t.Fatalf("LoadConfig(%q) with %s=%s = %+v, nil - want an error", prefix, tt.key, tt.value, got)
			}
			if !strings.Contains(err.Error(), tt.key) {
				t.Errorf("LoadConfig(%q) with %s=%s: error %q does not mention the variable name %s",
					prefix, tt.key, tt.value, err, tt.key)
			}
		})
	}

	// A different prefix reads different variables.
	t.Setenv("GBE_OTHERCFG_HOST", "other.example.com")
	unsetenv(t, "GBE_OTHERCFG_PORT")
	unsetenv(t, "GBE_OTHERCFG_DEBUG")
	want := Config{Host: "other.example.com", Port: 8080}
	if got, err := LoadConfig("GBE_OTHERCFG_"); err != nil || got != want {
		t.Errorf(`LoadConfig("GBE_OTHERCFG_") = %+v, %v, want %+v, nil`, got, err, want)
	}
}
