package cmd

import (
	"bufio"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Deletes the specified project folder.",
	Long: `The 'clean' command deletes the specified project folder and all of its contents.
It requires confirmation before proceeding with the deletion to avoid accidental removal of important data.`,
	RunE: CleanFunc,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func CleanFunc(cmd *cobra.Command, args []string) error {
	rootPath, err := currentDirectory(args)
	if err != nil {
		return err
	}

	projectPath := filepath.Join(rootPath, internal.RpmBuildName)

	if _, err := os.Stat(projectPath); err == nil {
		if askForConfirmation(rootPath) {
			if err := os.RemoveAll(projectPath); err != nil {
				return err
			}
			fmt.Printf("%s deleted successfully.\n", projectPath)
		}
	}

	return nil
}

func askForConfirmation(rootPath string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure you want to delete the rpm build project? type '%s' to confirm: ", filepath.Base(rootPath))

	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return askForConfirmation(rootPath)
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == filepath.Base(rootPath) {
		return true
	}

	fmt.Println("Invalid input. Please try again.")
	return askForConfirmation(rootPath)
}
