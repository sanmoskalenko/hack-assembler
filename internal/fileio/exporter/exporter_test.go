package exporter_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sanmoskalenko/hack-assembler/internal/fileio/exporter"
)

func TestExportFileWithContent(t *testing.T) {
	dir := t.TempDir()
	fileName := "Test.hack"
	content := []byte("1110000000000000")

	err := exporter.Export(fileName, dir, content)

	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("File content mismatch: got %q, want %q", data, content)
	}
}

func TestExportFileWithContentWithCreatingDirectory(t *testing.T) {
	parent := t.TempDir()
	subdir := filepath.Join(parent, "nested/dir")
	fileName := "output.hack"
	content := []byte("1010101010101010")

	err := exporter.Export(fileName, subdir, content)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	fullPath := filepath.Join(subdir, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Errorf("File was not created at expected location: %s", fullPath)
	}
}
