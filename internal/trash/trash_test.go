package trash

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"soft-rm/pkg/config"
	"testing"

	"github.com/spf13/viper"
)

func TestMoveToTrash(t *testing.T) {
	// Setup a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test-trash")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a dummy file to move
	dummyFile := filepath.Join(tmpDir, "dummy.txt")
	if err := ioutil.WriteFile(dummyFile, []byte("hello"), 0666); err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}

	// Setup a temporary trash directory
	trashDir := filepath.Join(tmpDir, "trash")
	viper.Set("trash_path", trashDir)
	viper.Set("retention_days", 30)
	config.SaveConfig()

	// Move the file to trash
	if err := MoveToTrash(dummyFile); err != nil {
		t.Fatalf("MoveToTrash failed: %v", err)
	}

	// Check if the file exists in the trash
	files, err := ioutil.ReadDir(trashDir)
	if err != nil {
		t.Fatalf("Failed to read trash dir: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("Expected 1 file in trash, got %d", len(files))
	}
}
