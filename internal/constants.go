package internal

import (
	"path/filepath"
)

const (
	RpmBuildName     = ".rpm"
	PackageName      = "package"
	RepoDataName     = "repodata"
	RepoDataFileName = "repodata.toml"
)

var (
	RepoDataPath  = filepath.Join(RpmBuildName, RepoDataName)
	BuildPath     = filepath.Join(RpmBuildName, PackageName, "BUILD")
	BuildRootPath = filepath.Join(RpmBuildName, PackageName, "BUILDROOT")
	RpmsPath      = filepath.Join(RpmBuildName, PackageName, "RPMS")
	SourcesPath   = filepath.Join(RpmBuildName, PackageName, "SOURCES")
	SpecsPath     = filepath.Join(RpmBuildName, PackageName, "SPECS")
	SrpmsPath     = filepath.Join(RpmBuildName, PackageName, "SRPMS")

	RepoDatabaseFile = filepath.Join(RpmBuildName, RepoDataName, "config.sqlite3")
	RepoDataFile     = filepath.Join(RpmBuildName, RepoDataName, RepoDataFileName)
)
