//go:build solution

package environmentvariables

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

func GetOrDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

func Status(key string) string {
	value, ok := os.LookupEnv(key)
	switch {
	case !ok:
		return key + " is not set"
	case value == "":
		return key + " is empty"
	default:
		return key + "=" + value
	}
}

func SetAll(vars map[string]string) error {
	for key, value := range vars {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}
	return nil
}

func EnvMap(prefix string) map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 && strings.HasPrefix(pair[0], prefix) {
			vars[pair[0]] = pair[1]
		}
	}
	return vars
}

func LoadConfig(prefix string) (Config, error) {
	cfg := Config{Host: "localhost", Port: 8080}

	if host := os.Getenv(prefix + "HOST"); host != "" {
		cfg.Host = host
	}
	if port := os.Getenv(prefix + "PORT"); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil {
			return Config{}, fmt.Errorf("%sPORT: %w", prefix, err)
		}
		cfg.Port = n
	}
	if debug := os.Getenv(prefix + "DEBUG"); debug != "" {
		b, err := strconv.ParseBool(debug)
		if err != nil {
			return Config{}, fmt.Errorf("%sDEBUG: %w", prefix, err)
		}
		cfg.Debug = b
	}
	return cfg, nil
}
