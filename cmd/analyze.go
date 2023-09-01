package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analysis command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: analyzeFunc,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}

func analyzeFunc(cmd *cobra.Command, args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = filepath.Join(wd, args[0])
		}
	}

	fmt.Println("analyze called")
	//TODO implement analyze function to analyze the project and generate the report to sqlite database
	//TODO check for zip, tar, 7z, cpio, xz and gzip files
	//TODO check for spec file
	//TODO check for source files
	//TODO check for patches
	//TODO check for changelog
	return nil
}
