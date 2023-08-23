package internal

import (
	"path/filepath"
)

const (
	RepoDataName  = ".repodata"
	GpgKeysName   = "gpgkeys"
	RpmBuildName  = "rpmbuild"
	BuildName     = "BUILD"
	BuildRootName = "BUILDROOT"
	RpmsName      = "RPMS"
	SourcesName   = "SOURCES"
	SpecsName     = "SPECS"
	SrpmsName     = "SRPMS"
)

var (
	GpgKeysPath = filepath.Join(RepoDataName, GpgKeysName)

	BuildPath     = filepath.Join(RpmBuildName, BuildName)
	BuildRootPath = filepath.Join(RpmBuildName, BuildRootName)
	RpmsPath      = filepath.Join(RpmBuildName, RpmsName)
	SourcesPath   = filepath.Join(RpmBuildName, SourcesName)
	SpecsPath     = filepath.Join(RpmBuildName, SpecsName)
	SrpmsPath     = filepath.Join(RpmBuildName, SrpmsName)

	RepoDatabaseFile = filepath.Join(RepoDataName, "config.sqlite3")
	RepoDataFile     = filepath.Join(RepoDataName, "repodata.toml")
)
