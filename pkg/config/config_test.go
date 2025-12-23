package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	// Setup a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test-config")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set the config path to the temp directory
	configPath = filepath.Join(tmpDir, "config.json")

	// Test creating a new config
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if cfg.RetentionDays != 30 {
		t.Fatalf("Expected default retention days to be 30, got %d", cfg.RetentionDays)
	}

	// Test loading an existing config
	viper.Set("retention_days", 60)
	SaveConfig()
	cfg, err = LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if cfg.RetentionDays != 60 {
		t.Fatalf("Expected retention days to be 60, got %d", cfg.RetentionDays)
	}
}
