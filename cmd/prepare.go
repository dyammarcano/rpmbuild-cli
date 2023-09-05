package cmd

import (
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/components"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"

	"github.com/spf13/cobra"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare package structure for building",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: prepareFunc,
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}

func prepareFunc(_ *cobra.Command, args []string) error {
	rootPath, err := directory.CurrentDirectory(args)
	if err != nil {
		return err
	}

	if err := components.Prepare(rootPath); err != nil {
		return err
	}

	fmt.Println("prepare called")
	//TODO source folder and make sure it is not empty then copy the source to the build folder
	//TODO check database for posible version and spec
	return nil
}
