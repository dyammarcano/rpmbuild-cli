package git

import (
	"encoding/json"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"path/filepath"
	"testing"
)

func TestChangelog(t *testing.T) {
	rootPath, err := directory.CurrentDirectory(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	rootPath = filepath.Join(rootPath, "..", "..") // go up two levels to get the root path

	changelogs, err := Changelog(rootPath)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(changelogs) == 0 {
		t.Errorf("Expected changelogs, got %v", changelogs)
	}

	jsonData, err := json.Marshal(changelogs)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	fmt.Println(string(jsonData))
}
