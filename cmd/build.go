package cmd

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/spf13/cobra"
	"os"
	"reflect"
	"sort"
	"strings"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: BuildFunc,
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func BuildFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("build called")

	if _, err := os.Stat(internal.RepoDataFile); err != nil {
		return err
	}

	var conf structures.Config
	meta, err := toml.DecodeFile(internal.RepoDataFile, &conf)
	if err != nil {
		return err
	}

	indent := strings.Repeat(" ", 14)

	fmt.Print("Decoded")
	typ, val := reflect.TypeOf(conf), reflect.ValueOf(conf)
	for i := 0; i < typ.NumField(); i++ {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 7)
		}
		fmt.Printf("%s%-11s → %v\n", indent, typ.Field(i).Name, val.Field(i).Interface())
	}

	fmt.Print("\nKeys")
	keys := meta.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	for i, k := range keys {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 10)
		}
		fmt.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	}

	fmt.Print("\nUndecoded")
	keys = meta.Undecoded()
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	for i, k := range keys {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 5)
		}
		fmt.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	}

	return nil
}
