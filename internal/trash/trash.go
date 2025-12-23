package trash

import (
	"fmt"
	"os"
	"path/filepath"
	"soft-rm/pkg/config"
	"time"
)

func MoveToTrash(path string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}

	trashPath := cfg.TrashPath
	if _, err := os.Stat(trashPath); os.IsNotExist(err) {
		if err := os.MkdirAll(trashPath, 0755); err != nil {
			return fmt.Errorf("could not create trash directory: %w", err)
		}
	}

	timestamp := time.Now().Format("20060102150405")
	baseName := filepath.Base(path)
	newPath := filepath.Join(trashPath, fmt.Sprintf("%s_%s", timestamp, baseName))

	return os.Rename(path, newPath)
}
