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
CriateFoldersStructure
└── rpm
    ├── package
    │	└── CriateFoldersStructureLinux
    └── repodata
        └── config.sqlite3
        └── repodata.toml
*/

// CriateFoldersStructure creates the folders structure for the RPM build
func CriateFoldersStructure(basePath string) error {
	directories := []string{
		internal.RepoDataPath,
		internal.PackagePath,
	}

	for _, directory := range directories {
		if err := os.MkdirAll(filepath.Join(basePath, directory), 0755); err != nil {
			return err
		}
	}

	if err := CriateFoldersStructureLinux(filepath.Join(basePath, internal.PackagePath)); err != nil {
		return err
	}

	return nil
}

/*
├── boot				(bootable kernel images, etc.)
├── dev					(device files)
├── etc					(host-specific system configuration)
│	└── systemd
│	 	└── system		(systemd unit files)
├── home				(user home directories)
├── lib					(shared libraries)
│	├── lib				(64-bit shared libraries)
│	└── systemd
│	 	└── system		(systemd unit files)
├── lib64				(64-bit shared libraries)
├── media				(removable media)
├── mnt					(mount point for a temporarily mounted filesystem)
├── opt					(add-on application software packages)
│	└── bin				(binaries for add-on packages)
├── proc				(process information pseudo-filesystem)
├── root				(root home directory)
├── run 				(run-time variable data)
├── sbin				(system binaries)
├── srv					(data for services provided by this system)
├── sys					(system information)
├── tmp					(temporary files)
├── usr					(read-only user data)
│	├── bin				(user binaries)
│	├── include			(C header files)
│	├── lib				(user libraries)
│	├── local			(secondary hierarchy)
│	├── sbin			(system binaries)
│	├── share			(architecture-independent data)
│	├── src				(source code)
│	└── tmp				(temporary files)
└── var					(variable data)
    ├── lib				(variable state information)
    ├── tmp				(temporary files)
    ├── cache			(application cache data)
    └── log				(log files)
*/

func CriateFoldersStructureLinux(basePath string) error {
	directories := []string{
		"boot",
		"dev",
		"etc",
		"etc/systemd",
		"etc/systemd/system",
		"home",
		"lib",
		"lib/lib",
		"lib/systemd",
		"lib/systemd/system",
		"lib64",
		"media",
		"mnt",
		"opt",
		"opt/bin",
		"proc",
		"root",
		"run",
		"sbin",
		"srv",
		"sys",
		"tmp",
		"usr",
		"usr/bin",
		"usr/include",
		"usr/lib",
		"usr/local",
		"usr/sbin",
		"usr/share",
		"usr/src",
		"usr/tmp",
		"var",
		"var/lib",
		"var/tmp",
		"var/cache",
		"var/log",
	}

	for _, directory := range directories {
		if err := os.MkdirAll(filepath.Join(basePath, directory), 0755); err != nil {
			panic(err)
		}
	}

	return nil
}

func CleanDirectoriesNotUsed(basePath string) error {
	var dirs []string

	if err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != basePath && d.IsDir() {
			dirs = append(dirs, path)
		}

		return nil
	}); err != nil {
		return err
	}

	for _, dir := range dirs {
		if err := os.Remove(dir); err != nil {
			return err
		}
	}

	return nil
}

func CurrentDirectory(args []string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = filepath.Join(wd, args[0])
		}
	}

	return wd, nil
}
