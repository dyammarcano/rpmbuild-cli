package cmd

import (
	"github.com/dyammarcano/rpmbuild-cli/internal"
	"github.com/dyammarcano/rpmbuild-cli/internal/directory"
	"github.com/spf13/cobra"
	"path/filepath"
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
	RunE: buildFunc,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildFunc(cmd *cobra.Command, args []string) error {
	wd, err := directory.CurrentDirectory(args)
	if err != nil {
		return err
	}

	_ = filepath.Join(wd, internal.RepoDataFile)

	//	//TODO helper decode encode toml file
	//
	//	if _, err := os.Stat(repodata); err != nil {
	//		return err
	//	}
	//
	//	var conf structures.Config
	//	meta, err := toml.DecodeFile(repodata, &conf)
	//	if err != nil {
	//		return err
	//	}
	//
	//	indent := strings.Repeat(" ", 14)
	//
	//	fmt.Print("Decoded")
	//	typ, val := reflect.TypeOf(conf), reflect.ValueOf(conf)
	//	for i := 0; i < typ.NumField(); i++ {
	//		indent := indent
	//		if i == 0 {
	//			indent = strings.Repeat(" ", 7)
	//		}
	//		fmt.Printf("%s%-11s → %v\n", indent, typ.Field(i).Name, val.Field(i).Interface())
	//	}
	//
	//	fmt.Print("\nKeys")
	//	keys := meta.Keys()
	//	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	//	for i, k := range keys {
	//		indent := indent
	//		if i == 0 {
	//			indent = strings.Repeat(" ", 10)
	//		}
	//		fmt.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	//	}
	//
	//	fmt.Print("\nUndecoded")
	//	keys = meta.Undecoded()
	//	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	//	for i, k := range keys {
	//		indent := indent
	//		if i == 0 {
	//			indent = strings.Repeat(" ", 5)
	//		}
	//		fmt.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	//	}
	//
	//TODO:
	//convert json to spec

	return nil
}
