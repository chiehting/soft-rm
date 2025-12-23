package integration

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binaryName = "soft-rm"
)

func TestMain(m *testing.M) {
	// Build the binary for testing
	cmd := exec.Command("go", "build", "-o", binaryName, "../cmd/soft-rm")
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// Run the tests
	exitCode := m.Run()

	// Clean up the binary
	os.Remove(binaryName)

	os.Exit(exitCode)
}

func TestCli(t *testing.T) {
	// Setup a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test-cli")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a dummy file to move
	dummyFile := filepath.Join(tmpDir, "dummy.txt")
	if err := ioutil.WriteFile(dummyFile, []byte("hello"), 0666); err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}

	// Run the soft-rm command
	cmd := exec.Command("./"+binaryName, dummyFile)
	if err := cmd.Run(); err != nil {
		t.Fatalf("soft-rm command failed: %v", err)
	}

	// Check if the file exists in the trash (we can't easily know the trash path, so we can't test this here)
	// This is a limitation of this test, a more advanced test could parse the config
}
