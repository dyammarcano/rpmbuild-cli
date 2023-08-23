package database

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"testing"
	"time"
)

func TestCreateDatabaseFile(t *testing.T) {
	db, err := NewDatabase("test.db")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	defer db.Close()

	if db == nil {
		t.Errorf("Expected a database connection, got %v", db)
	}

	if err := db.Migrate(
		&structures.Package{},
		&structures.PackageFile{},
		&structures.PackageVersion{},
		&structures.PackageProvide{},
		&structures.PackageRequire{},
		&structures.Changelog{},
		&structures.Spec{},
		&structures.File{},
	); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.Package{
		Name: "test",
		Arch: "x86_64",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.PackageFile{
		PkgID: "test",
		File:  "test",
		Type:  "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.PackageVersion{
		PkgID:   "test",
		Epoch:   "test",
		Version: "test",
		Release: "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.PackageProvide{
		PkgID:   "test",
		Name:    "test",
		Version: "test",
		Release: "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.PackageRequire{
		PkgID:   "test",
		Name:    "test",
		Version: "test",
		Release: "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.Changelog{
		Author:    "test",
		Date:      time.Now().String(),
		Changelog: "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.Spec{
		Name:             "test",
		Version:          "test",
		Release:          "test",
		Epoch:            "test",
		Arch:             "test",
		Summary:          "test",
		Description:      "test",
		URL:              "test",
		License:          "test",
		Group:            nil,
		BuildHost:        "test",
		BuildTime:        "test",
		Source:           "test",
		BuildRequires:    "test",
		Requires:         nil,
		Provides:         nil,
		Conflicts:        nil,
		Obsoletes:        nil,
		Changelog:        "test",
		Files:            "test",
		ChangelogTime:    "test",
		ChangelogName:    "test",
		ChangelogText:    "test",
		ChangelogEmail:   "test",
		ChangelogVersion: "test",
		ChangelogRelease: "test",
		ChangelogOrder:   "test",
		ChangelogFlags:   "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := db.Create(&structures.File{
		Name:     "test",
		Size:     1,
		Mode:     nil,
		MD5:      nil,
		SHA1:     nil,
		SHA256:   nil,
		SHA512:   nil,
		RMD160:   nil,
		LinkTo:   "test",
		Flags:    "test",
		Username: "test",
	}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
