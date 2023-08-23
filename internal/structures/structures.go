package structures

import "time"

type (
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

	Changelog struct {
		PkgKey    int `gorm:"primaryKey"`
		Author    string
		Date      string
		Changelog string
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
)
