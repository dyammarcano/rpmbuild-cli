package directory

import (
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type (
	FileMode uint32

	DirEntry struct {
		FilePath     string
		IsDir        bool
		Type         FileMode
		Size         int64
		ModeFileMode fs.FileMode
		ModTime      time.Time
		Sys          any
	}

	Info struct {
		FileInfo fs.FileInfo
	}
)

func WalkEntries(path string, dirs bool) ([]DirEntry, error) {
	if dirs {
		return walkDirEntry(path)
	}

	return walkFiles(path)
}

func addEntry(filePath string, isDir bool, mode FileMode, size int64, modTime time.Time, sys interface{}) DirEntry {
	return DirEntry{
		FilePath:     filePath,
		IsDir:        isDir,
		Type:         mode,
		Size:         size,
		ModeFileMode: fs.FileMode(mode),
		ModTime:      modTime,
		Sys:          sys,
	}
}

func walkDirEntry(path string) ([]DirEntry, error) {
	var entries []DirEntry
	err := filepath.WalkDir(path, func(p string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			entries = append(entries, addEntry(p, info.IsDir(), FileMode(info.Type()), 0, time.Time{}, nil))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func walkFiles(path string) ([]DirEntry, error) {
	var entries []DirEntry
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			entries = append(entries, addEntry(p, info.IsDir(), FileMode(info.Mode()), info.Size(), info.ModTime(), info.Sys()))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}

/*
├── rpmbuild
│   ├── BUILD
│   ├── BUILDROOT
│   ├── RPMS
│   ├── SOURCES
│   ├── SPECS
│   └── SRPMS
└── .repodata
    └── gpgkeys
*/

// CriateFoldersStructure creates the folders structure for the RPM build
func CriateFoldersStructure(basePath string) {
	if basePath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		basePath = cwd
	}

	directories := []string{
		internal.RepoDataName,
		internal.GpgKeysPath,
		internal.BuildPath,
		internal.BuildRootPath,
		internal.RpmsPath,
		internal.SourcesPath,
		internal.SpecsPath,
		internal.SrpmsPath,
	}

	for _, directory := range directories {
		if err := os.MkdirAll(filepath.Join(basePath, directory), 0755); err != nil {
			panic(err)
		}
	}
}
