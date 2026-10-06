//go:build solution

package embeddirective

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed assets/motd.txt
var motd string

//go:embed assets/motd.txt
var motdBytes []byte

//go:embed assets/config/*.env
var configs embed.FS

func ConfigNames() ([]string, error) {
	// ReadDir returns the entries sorted by file name.
	entries, err := configs.ReadDir("assets/config")
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}

func ConfigValue(env, key string) (string, error) {
	// Embedded paths use forward slashes even on Windows, so plain string
	// concatenation (or path.Join) is right here - not filepath.Join.
	data, err := configs.ReadFile("assets/config/" + env + ".env")
	if err != nil {
		return "", err
	}
	for line := range strings.Lines(string(data)) {
		k, v, ok := strings.Cut(strings.TrimSpace(line), "=")
		if ok && k == key {
			return v, nil
		}
	}
	return "", fmt.Errorf("key %q not found in %s config", key, env)
}
