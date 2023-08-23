package initialfile

import (
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

func InitialFile(basePath string) bool {
	data, err := toml.Marshal(structures.Config{})
	if err != nil {
		return false
	}

	filePath := filepath.Join(basePath, internal.RepoDataFile)

	f, err := os.Create(filePath)
	if err != nil {
		return false
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			return
		}
	}(f)

	if _, err = f.Write(data); err != nil {
		return false
	}

	return true
}
