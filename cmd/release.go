package cmd

import (
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: releaseFunc,
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}

func releaseFunc(cmd *cobra.Command, args []string) error {
	_, err := directory.CurrentDirectory(args)
	if err != nil {
		return err
	}

	fmt.Println("release called")
	//TODO check the database for the source and spec file
	//TODO check the database for the version
	//TODO check the database for the build folder
	return nil
}
