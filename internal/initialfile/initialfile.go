package initialfile

import (
	"encoding/json"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/git"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func InitialFile(basePath string) bool {
	data, err := toml.Marshal(structures.Config{})
	if err != nil {
		return false
	}

	f, err := os.Create(filepath.Join(basePath, internal.RepoDataFile))
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

func InitialMetadataFile(basePath string, kind structures.MetadataType) error {
	switch kind {
	case structures.MetadataTypeToml:
		return createMetadataFile(basePath, internal.MetadataTomlFileName, toml.Marshal)
	case structures.MetadataTypeJson:
		return createMetadataFile(basePath, internal.MetadataJsonFileName, json.Marshal)
	case structures.MetadataTypeYaml:
		return createMetadataFile(basePath, internal.MetadataYamlFileName, yaml.Marshal)
	}
	return nil
}

type marshalFunc func(v any) ([]byte, error)

func createMetadataFile(rootPath, fileName string, marshal marshalFunc) error {
	rootPath = filepath.Join(rootPath, "..")

	info, err := git.GetGitInfo(rootPath)
	if err != nil {
		return err
	}

	metadata := &structures.Metadata{
		Name:          filepath.Base(rootPath),
		Version:       info.Version,
		Commit:        info.CommitHash,
		Summary:       "This is a test package",
		Description:   "This is a test package",
		License:       structures.LicenseProprietary,
		Group:         "Development/Tools",
		URL:           info.RemoteRepo,
		BuildHost:     structures.BuildHostLinux,
		BuildArch:     structures.BuildArchX64,
		Source:        "https://github.com/dyammarcano/utils/archive/refs/tags/v0.1.3.tar.gz",
		BuildRequires: "go >= 1.20",
		Requires:      "go >= 1.20",
		Provides:      "utils",
		Conflicts:     "utils",
		Obsoletes:     "utils",
		Changelog:     info.Changelog,
	}

	data, err := marshal(metadata)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(rootPath, internal.MetadataFile))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			return
		}
	}(f)

	_, err = f.Write(data)
	return err
}
