package cleaner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"soft-rm/pkg/config"
	"syscall"
	"time"
)

func SpawnCleanupProcess() {
	go func() {
		cfg, err := config.LoadConfig()
		if err != nil {
			// Cannot do much here, maybe log to a file in the future
			return
		}

		lockFilePath := filepath.Join(cfg.TrashPath, ".lock")
		lockFile, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return
		}
		defer lockFile.Close()

		if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
			// Another process is already cleaning up
			return
		}
		defer syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)

		files, err := ioutil.ReadDir(cfg.TrashPath)
		if err != nil {
			return
		}

		retentionDuration := time.Duration(cfg.RetentionDays) * 24 * time.Hour
		now := time.Now()

		for _, file := range files {
			if file.Name() == ".lock" {
				continue
			}

			if now.Sub(file.ModTime()) > retentionDuration {
				filePath := filepath.Join(cfg.TrashPath, file.Name())
				if err := os.RemoveAll(filePath); err != nil {
					// Log error
				}
			}
		}
	}()
}
