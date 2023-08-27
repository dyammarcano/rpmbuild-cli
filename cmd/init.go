package cmd

import (
	"errors"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/database"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"github.com/dyammarcano/rpmbuild-cli/internal/initialfile"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/dyammarcano/utils/display"
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

func InitFunc(_ *cobra.Command, args []string) error {
	rootPath, err := currentDirectory(args)
	if err != nil {
		return err
	}

	projectPath := filepath.Join(rootPath, internal.RpmBuildName)

	if _, err := os.Stat(projectPath); err == nil {
		return errors.New(fmt.Sprintf("%s already exists", projectPath))
	}

	if !checkIfGitInitialized(rootPath) {
		return errors.New("the current directory not have git initialized")
	}

	if err := directory.CriateFoldersStructure(rootPath); err != nil {
		return err
	}
	fmt.Println("* package structure created")

	if !initialfile.InitialFile(rootPath) {
		return errors.New("failed to create initial file")
	}
	fmt.Printf("* %s created\n", internal.RepoDataFileName)

	db, err := database.NewDatabase(filepath.Join(rootPath, internal.RepoDatabaseFile))
	if err != nil {
		return err
	}
	fmt.Println("* report database created")

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
	fmt.Println("* database migration successful")

	if err := display.DisplayDirectoryTree(rootPath); err != nil {
		fmt.Println("Error:", err)
	}

	return nil
}

func checkIfGitInitialized(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		return true
	}
	return false
}
