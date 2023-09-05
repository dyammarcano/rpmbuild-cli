package structures

import "time"

const (
	MetadataTypeToml MetadataType = "toml"
	MetadataTypeJson MetadataType = "json"
	MetadataTypeYaml MetadataType = "yaml"

	BuildHostLinux   BuildHost = "linux"
	BuildHostWindows BuildHost = "windows"

	BuildArchX64     BuildArch = "x86_64"
	BuildArchI386    BuildArch = "i386"
	BuildArchI686    BuildArch = "i686"
	BuildArchAarch64 BuildArch = "aarch64"

	LicenseMIT         License = "MIT"
	LicenseGPL         License = "GPL"
	LicenseApache      License = "Apache"
	LicenseBSD         License = "BSD"
	LicenseLGPL        License = "LGPL"
	LicensePublic      License = "Public"
	LicenseProprietary License = "Proprietary"
)

type (
	MetadataType string
	BuildHost    string
	BuildArch    string
	License      string

	Metadata struct {
		Name          string      `json:"name" yaml:"name" toml:"name"`
		Version       string      `json:"version" yaml:"version" toml:"version"`
		Commit        string      `json:"commit" yaml:"commit" toml:"commit"`
		Epoch         string      `json:"epoch" yaml:"epoch" toml:"epoch"`
		Arch          string      `json:"arch" yaml:"arch" toml:"arch"`
		Summary       string      `json:"summary" yaml:"summary" toml:"summary"`
		Description   string      `json:"description" yaml:"description" toml:"description"`
		URL           string      `json:"url" yaml:"url" toml:"url"`
		License       License     `json:"license" yaml:"license" toml:"license"`
		Group         string      `json:"group" yaml:"group" toml:"group"`
		BuildHost     BuildHost   `json:"buildhost" yaml:"buildhost" toml:"buildhost"`
		BuildArch     BuildArch   `json:"buildarch" yaml:"buildarch" toml:"buildarch"`
		Source        string      `json:"source" yaml:"source" toml:"source"`
		BuildRequires string      `json:"buildrequires" yaml:"buildrequires" toml:"buildrequires"`
		Requires      string      `json:"requires" yaml:"requires" toml:"requires"`
		Provides      string      `json:"provides" yaml:"provides" toml:"provides"`
		Conflicts     string      `json:"conflicts" yaml:"conflicts" toml:"conflicts"`
		Obsoletes     string      `json:"obsoletes" yaml:"obsoletes" toml:"obsoletes"`
		Changelog     []Changelog `json:"changelog" yaml:"changelog" toml:"changelog"`
		//Files         string    `json:"files" yaml:"files" toml:"files"`
	}

	Config struct {
		Host       string    `toml:"host"`
		User       string    `toml:"user"`
		Pass       string    `toml:"pass"`
		Db         string    `toml:"db"`
		Age        int       `toml:"age"`
		Cats       []string  `toml:"cats"`
		Pi         float64   `toml:"pi"`
		Perfection []int     `toml:"perfection"`
		DOB        time.Time `toml:"dob"`
	}

	Changelog struct {
		Hash    string `gorm:"primaryKey" json:"hash"`
		Author  string `json:"author"`
		Email   string `json:"email"`
		Date    string `json:"date"`
		Title   string `json:"title"`
		Message string `json:"message"`
	}

	Package struct {
		PkgID     string `gorm:"primarykey"`
		Name      string
		Arch      string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	PackageFile struct {
		PkgID     string `gorm:"primaryKey"`
		File      string
		Type      string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoCreateTime"`
	}

	PackageVersion struct {
		PkgID     string `gorm:"primaryKey"`
		Epoch     string
		Version   string
		Release   string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoCreateTime"`
		Package   Package   `gorm:"foreignKey:PkgID"`
	}

	PackageProvide struct {
		PkgID     string `gorm:"primaryKey"`
		Name      string
		Version   string
		Release   string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoCreateTime"`
	}

	PackageRequire struct {
		PkgID     string `gorm:"primaryKey"`
		Name      string
		Version   string
		Release   string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoCreateTime"`
	}

	Spec struct {
		ID               int `gorm:"primaryKey"`
		Name             string
		Version          string
		Release          string
		Epoch            string
		Arch             string
		Summary          string
		Description      string
		URL              string
		License          string
		Group            *string
		BuildHost        string
		BuildTime        string
		Source           string
		BuildRequires    string
		Requires         *string
		Provides         *string
		Conflicts        *string
		Obsoletes        *string
		Changelog        string
		Files            string
		ChangelogTime    string
		ChangelogName    string
		ChangelogText    string
		ChangelogEmail   string
		ChangelogVersion string
		ChangelogRelease string
		ChangelogOrder   string
		ChangelogFlags   string
		CreatedAt        time.Time `gorm:"autoCreateTime"`
		UpdatedAt        time.Time `gorm:"autoCreateTime"`
	}

	File struct {
		ID        int `gorm:"primaryKey"`
		Name      string
		Size      int
		Mode      *string
		MD5       *string
		SHA1      *string
		SHA256    *string
		SHA512    *string
		RMD160    *string
		LinkTo    string
		Flags     string
		Username  string
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoCreateTime"`
	}

	FilePermission struct {
		Mode  string
		File  string
		Owner string
	}

	DirectoryPermission struct {
		Mode  string
		Path  string
		Owner string
	}
)
