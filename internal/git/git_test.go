package git

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"path/filepath"
	"testing"
)

func TestGetGitInfo(t *testing.T) {
	rootPath, err := directory.CurrentDirectory(nil)
	if err != nil {
		t.Error(err)
	}

	rootPath = filepath.Join(rootPath, "..", "..")

	g, err := GetGitInfo(rootPath)
	if err != nil {
		t.Error(err)
	}

	if g == nil {
		t.Error("git info is nil")
	}
}
