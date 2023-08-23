package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: CheckFunc,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func CheckFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("check called")
	//TODO check the database for the source and spec file
	//TODO check the database for the version
	//TODO check the database for the build folder
	return nil
}
