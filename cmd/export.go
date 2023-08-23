package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: ExportFunc,
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

func ExportFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("export called")
	//TODO dump all the data from the database to a file
	return nil
}
