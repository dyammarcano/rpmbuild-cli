package cmd

import (
	"bufio"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Deletes the RPM build project in the current directory.",
	Long: `The 'clean' command deletes the specified project folder and all of its contents.
It requires confirmation before proceeding with the deletion to avoid accidental removal of important data.`,
	RunE: cleanFunc,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func cleanFunc(cmd *cobra.Command, args []string) error {
	rootPath, err := directory.CurrentDirectory(args)
	if err != nil {
		return err
	}

	projectPath := filepath.Join(rootPath, internal.RpmBuildName)

	if _, err := os.Stat(projectPath); err == nil {
		if askForConfirmationClean(rootPath) {
			if err := os.RemoveAll(projectPath); err != nil {
				return err
			}
			fmt.Printf("%s deleted successfully.\n", projectPath)
		}
	}

	return nil
}

func askForConfirmationClean(rootPath string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure you want to delete the .rpm build project? type '%s' to confirm: ", filepath.Base(rootPath))

	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	if strings.TrimSpace(response) == filepath.Base(rootPath) {
		return true
	}

	fmt.Println("Invalid input. Please try again.")
	return askForConfirmationClean(rootPath)
}
