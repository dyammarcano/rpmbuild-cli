package cmd

import (
	"bufio"
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
	"strings"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the RPM build project in the current directory.",
	Long: `This 'init' command prepares the current directory for a RPM build project. 
It creates necessary folders structures, initializes files and sets up a database for the project.
It also checks if git has been initialized in the directory.

Use it like this: 

rpmbuild-cli init

Remember, you need to confirm the operation when asked "Are you sure you want to create the .rpm build project?", by typing 'yes' or 'no'.`,
	RunE: initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initFunc(_ *cobra.Command, args []string) error {
	rootPath, err := currentDirectory(args)
	if err != nil {
		return err
	}

	if askForConfirmationCreate(rootPath) {
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

		if err := createDatabase(rootPath); err != nil {
			return err
		}

		if err := display.DisplayDirectoryTree(rootPath); err != nil {
			fmt.Println("Error:", err)
		}
	}

	return nil
}

func createDatabase(rootPath string) error {
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

	return nil
}

func checkIfGitInitialized(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		return true
	}
	return false
}

func askForConfirmationCreate(rootPath string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure you want to create the .rpm build project? type 'yes' or 'no': ")

	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "yes" || response == "y" {
		return true
	}

	if response == "no" || response == "n" {
		return false
	}

	fmt.Println("Invalid input. Please try again.")
	return askForConfirmationCreate(rootPath)
}
