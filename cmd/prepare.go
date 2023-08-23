package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: PrepareFunc,
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}

func PrepareFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("prepare called")
	//TODO source folder and make sure it is not empty then copy the source to the build folder
	//TODO check database for posible version and spec
	return nil
}
