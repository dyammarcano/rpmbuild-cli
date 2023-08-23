package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: CleanFunc,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func CleanFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("clean called")
	//TODO delete the build folder
	return nil
}
