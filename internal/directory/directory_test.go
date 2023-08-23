package directory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TempDirectoryWithFiles(t *testing.T) (dirPath string, cleanup func()) {
	t.Helper()

	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}

	// make second level directory
	dir2 := filepath.Join(dir, "dir2")
	if err := os.Mkdir(dir2, 0755); err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}

	// make third level directory
	dir3 := filepath.Join(dir2, "dir3")
	if err := os.Mkdir(dir3, 0755); err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}

	files := []string{
		"file1.txt",
		"file2.txt",
	}

	for _, file := range files {
		path := filepath.Join(dir, file)
		f, err := os.Create(path)
		if err != nil {
			t.Fatalf("failed to create temp file: %v, %v", file, err)
		}

		if _, err := f.Write([]byte("This is a test file")); err != nil {
			if err := f.Close(); err != nil {
				return "", nil
			}
			t.Fatalf("failed to write into the file: %v", err)
		}
		if err := f.Close(); err != nil {
			return "", nil
		}
	}

	return dir, func() {
		err := os.RemoveAll(dir)
		if err != nil {
			t.Fatalf("failed to cleanup temp directory: %v", err)
		}
	}
}

func TestDirectory_Walk(t *testing.T) {
	dir, cleanup := TempDirectoryWithFiles(t)
	defer cleanup()

	entries, err := WalkEntries(dir, false)
	expected := 2

	if len(entries) != expected {
		t.Errorf("Got %d entries, expected %d", len(entries), expected)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		t.Errorf("Failed to marshal entries: %v", err)
	}

	t.Logf("Entries: %s", data)
}

func TestDirectory_WalkDir(t *testing.T) {
	dir, cleanup := TempDirectoryWithFiles(t)
	defer cleanup()

	entries, err := WalkEntries(dir, true)
	expected := 3

	if len(entries) != expected {
		t.Errorf("Got %d entries, expected %d", len(entries), expected)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		t.Errorf("Failed to marshal entries: %v", err)
	}

	t.Logf("Entries: %s", data)
}
