package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	TrashPath     string `json:"trash_path"`
	RetentionDays int    `json:"retention_days"`
}

var (
	cfg         *Config
	configPath  = "~/.config/soft-rm/config.json"
	defaultPath = "~/.soft-rm"
)

func LoadConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	p, err := expandPath(configPath)
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(p)
	viper.SetConfigType("json")

	// Set default values
	trashPath, err := expandPath(defaultPath)
	if err != nil {
		return nil, err
	}
	viper.SetDefault("trash_path", trashPath)
	viper.SetDefault("retention_days", 30)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create it with defaults
			if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
				return nil, err
			}
			if err := viper.SafeWriteConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	cfg = &Config{
		TrashPath:     viper.GetString("trash_path"),
		RetentionDays: viper.GetInt("retention_days"),
	}

	// Ensure trashPath is always expanded after loading/defaulting
	expandedTrashPath, err := expandPath(cfg.TrashPath)
	if err != nil {
		return nil, err
	}
	cfg.TrashPath = expandedTrashPath

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

func SaveConfig() error {
	p, err := expandPath(configPath)
	if err != nil {
		return err
	}

	cfg.TrashPath = viper.GetString("trash_path")
	cfg.RetentionDays = viper.GetInt("retention_days")

	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}

func expandPath(path string) (string, error) {
	if path[:2] == "~/" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[2:])
	}
	return path, nil
}
