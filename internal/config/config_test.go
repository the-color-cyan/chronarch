package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDevConfigUsesWorkingDirectory(t *testing.T) {
	t.Setenv("CHRONARCH_DEV", "1")
	t.Chdir(t.TempDir())

	cfg, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := cfg.GetString("storage.path"), filepath.Join(".chronarch", "chronarch.db"); got != want {
		t.Fatalf("storage.path = %q, want %q", got, want)
	}

	path := filepath.Join(".chronarch", "config.toml")
	if err := os.WriteFile(path, []byte("[storage]\npath = \"custom.db\"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	cfg, err = New()
	if err != nil {
		t.Fatal(err)
	}
	if got := cfg.GetString("storage.path"); got != "custom.db" {
		t.Fatalf("storage.path = %q, want custom.db", got)
	}
}
