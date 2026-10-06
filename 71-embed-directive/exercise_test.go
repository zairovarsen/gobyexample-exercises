package embeddirective

import (
	"slices"
	"strings"
	"testing"
)

const wantMotd = "Welcome, gopher!"

func TestMotdString(t *testing.T) {
	if got := strings.TrimSpace(motd); got != wantMotd {
		t.Errorf("motd = %q, want the content of assets/motd.txt (%q)", got, wantMotd)
	}
}

func TestMotdBytes(t *testing.T) {
	if got := strings.TrimSpace(string(motdBytes)); got != wantMotd {
		t.Errorf("motdBytes = %q, want the content of assets/motd.txt (%q)", got, wantMotd)
	}
}

func TestConfigsFS(t *testing.T) {
	for _, name := range []string{"assets/config/dev.env", "assets/config/prod.env"} {
		data, err := configs.ReadFile(name)
		if err != nil {
			t.Errorf("configs.ReadFile(%q) failed: %v", name, err)
			continue
		}
		if !strings.Contains(string(data), "HOST=") {
			t.Errorf("configs.ReadFile(%q) = %q, want the content of the file", name, data)
		}
	}
	if _, err := configs.ReadFile("assets/config/notes.txt"); err == nil {
		t.Error("notes.txt is embedded too, want only the *.env files in configs")
	}
	if _, err := configs.ReadFile("assets/motd.txt"); err == nil {
		t.Error("motd.txt is embedded in configs, want only the *.env files")
	}
}

func TestConfigNames(t *testing.T) {
	got, err := ConfigNames()
	if err != nil {
		t.Fatalf("ConfigNames() returned an error: %v", err)
	}
	if want := []string{"dev.env", "prod.env"}; !slices.Equal(got, want) {
		t.Errorf("ConfigNames() = %v, want %v", got, want)
	}
}

func TestConfigValue(t *testing.T) {
	tests := []struct {
		env, key string
		want     string
	}{
		{"dev", "HOST", "localhost"},
		{"dev", "PORT", "8080"},
		{"dev", "DEBUG", "true"},
		{"prod", "HOST", "example.com"},
		{"prod", "PORT", "443"},
	}
	for _, tt := range tests {
		got, err := ConfigValue(tt.env, tt.key)
		if err != nil {
			t.Errorf("ConfigValue(%q, %q) returned an error: %v", tt.env, tt.key, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ConfigValue(%q, %q) = %q, want %q", tt.env, tt.key, got, tt.want)
		}
	}

	if got, err := ConfigValue("staging", "HOST"); err == nil {
		t.Errorf("ConfigValue(\"staging\", \"HOST\") = %q with no error, want an error: there is no staging.env", got)
	}
	if got, err := ConfigValue("dev", "PASSWORD"); err == nil {
		t.Errorf("ConfigValue(\"dev\", \"PASSWORD\") = %q with no error, want an error: the key does not exist", got)
	}
	if got, err := ConfigValue("dev", "POR"); err == nil {
		t.Errorf("ConfigValue(\"dev\", \"POR\") = %q with no error, want an error: keys must match exactly", got)
	}
}
