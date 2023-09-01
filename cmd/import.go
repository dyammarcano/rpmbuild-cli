package cmd

import (
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: importFunc,
}

func init() {
	rootCmd.AddCommand(importCmd)
}

func importFunc(cmd *cobra.Command, args []string) error {
	_, err := directory.CurrentDirectory(args)
	if err != nil {
		return err
	}

	fmt.Println("import called")
	//TODO import the data from the file to the database
	return nil
}
