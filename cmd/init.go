package cmd

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: InitFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("directory", "f", false, "Create the directory structure")

	if err := viper.BindPFlag("directory", initCmd.Flags().Lookup("directory")); err != nil {
		panic(err)
	}
}

func InitFunc(cmd *cobra.Command, args []string) error {
	directory.CriateFoldersStructure(viper.GetString("directory"))
	return nil
}
