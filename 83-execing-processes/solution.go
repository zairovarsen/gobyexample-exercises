//go:build solution

package execingprocesses

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"
	"syscall"
)

func FindBinary(name string) (string, error) {
	binary, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("find %s: %w", name, err)
	}
	return binary, nil
}

func Argv(name string, args ...string) []string {
	argv := make([]string, 0, len(args)+1)
	argv = append(argv, name)
	return append(argv, args...)
}

func EnvWith(env []string, key, value string) []string {
	result := slices.Clone(env)
	entry := key + "=" + value
	for i, kv := range result {
		if strings.HasPrefix(kv, key+"=") {
			result[i] = entry
			return result
		}
	}
	return append(result, entry)
}

func Replace(name string, args []string, env []string) error {
	binary, err := FindBinary(name)
	if err != nil {
		return err
	}
	return syscall.Exec(binary, Argv(name, args...), env)
}
