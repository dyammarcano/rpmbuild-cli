package cmd

import (
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/database"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"github.com/dyammarcano/rpmbuild-cli/internal/initialfile"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: InitFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func InitFunc(cmd *cobra.Command, args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = filepath.Join(wd, args[0])
		}
	}

	if err := directory.CriateFoldersStructure(wd); err != nil {
		return err
	}

	initialfile.InitialFile(wd)

	databaseFile := filepath.Join(wd, internal.RepoDatabaseFile)

	db, err := database.NewDatabase(databaseFile)
	if err != nil {
		return err
	}

	defer db.Close()

	if db == nil {
		return err
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
		return err
	}

	return nil
}
