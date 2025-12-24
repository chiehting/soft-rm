package cleaner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"soft-rm/pkg/config"
	"syscall"
	"time"
)

func SpawnCleanupProcess() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cleaner: error getting executable path: %v\n", err)
		return
	}

	cmd := exec.Command(executable, "--background-cleanup")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "cleaner: error starting background process: %v\n", err)
	}
}

func RunCleanup() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cleaner: error loading config: %v\n", err)
		return
	}

	lockFilePath := filepath.Join(cfg.TrashPath, ".lock")
	lockFile, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		// Can't open lock file, another process is likely running or there's a permissions issue.
		// Silently exit.
		return
	}
	defer lockFile.Close()

	if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		// Another process is already cleaning up.
		// Silently exit.
		return
	}
	defer syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)

	files, err := os.ReadDir(cfg.TrashPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cleaner: error reading trash directory: %v\n", err)
		return
	}

	retentionDuration := time.Duration(cfg.RetentionDays) * 24 * time.Hour
	now := time.Now()

	for _, file := range files {
		if file.Name() == ".lock" {
			continue
		}

		info, err := file.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cleaner: error getting file info for %s: %v\n", file.Name(), err)
			continue
		}

		age := now.Sub(info.ModTime())
		if age > retentionDuration {
			filePath := filepath.Join(cfg.TrashPath, file.Name())
			if err := os.RemoveAll(filePath); err != nil {
				fmt.Fprintf(os.Stderr, "cleaner: error deleting file %s: %v\n", filePath, err)
			}
		}
	}
}
