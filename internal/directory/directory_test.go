package directory

import (
	"encoding/json"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/utils/mocks"
	"os"
	"path/filepath"
	"testing"
)

func TempDirectoryWithFiles(t *testing.T) (dirPath string, cleanup func()) {
	t.Helper()

	dir, err := mocks.CreateTempDir(t)
	if err != nil {
		t.Errorf("Failed to create tmp dir: %v", err)
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
		f, err := os.Create(filepath.Join(dir, file))
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
		if err := os.RemoveAll(dir); err != nil {
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

func TestDirectory_CriateFoldersStructure(t *testing.T) {
	dir, err, cleanup := mocks.CreateTempDirCleanUp(t)
	if err != nil {
		t.Errorf("Failed to create tmp dir: %v", err)
	}
	defer cleanup()

	if err := CriateFoldersStructure(dir); err != nil {
		t.Errorf("Failed to create folders structure: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, internal.RpmsPath)); os.IsNotExist(err) {
		t.Errorf("package directory was not created")
	}
}
