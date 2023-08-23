package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

type (
	Config struct {
		Host       string    `toml:"host"`
		User       string    `toml:"user"`
		Pass       string    `toml:"pass"`
		Db         string    `toml:"db"`
		Age        int       `toml:"age"`
		Cats       []string  `toml:"cats"`
		Pi         float64   `toml:"pi"`
		Perfection []int     `toml:"perfection"`
		DOB        time.Time `toml:"dob"`
	}
)

const (
	RepoDataName  = ".repodata"
	GpgKeysName   = "gpgkeys"
	RpmBuildName  = "rpmbuild"
	BuildName     = "BUILD"
	BuildRootName = "BUILDROOT"
	RpmsName      = "RPMS"
	SourcesName   = "SOURCES"
	SpecsName     = "SPECS"
	SrpmsName     = "SRPMS"
)

var (
	GpgKeysPath = filepath.Join(RepoDataName, GpgKeysName)

	BuildPath     = filepath.Join(RpmBuildName, BuildName)
	BuildRootPath = filepath.Join(RpmBuildName, BuildRootName)
	RpmsPath      = filepath.Join(RpmBuildName, RpmsName)
	SourcesPath   = filepath.Join(RpmBuildName, SourcesName)
	SpecsPath     = filepath.Join(RpmBuildName, SpecsName)
	SrpmsPath     = filepath.Join(RpmBuildName, SrpmsName)

	RepoDatabaseFile = filepath.Join(RepoDataName, "config.sqlite3")
	RepoDataFile     = filepath.Join(RepoDataName, "repodata.toml")
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
	Run: Build,
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

func Build(cmd *cobra.Command, args []string) {
	fmt.Println("build called")

	if _, err := os.Stat(RepoDataFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var conf Config
	meta, err := toml.DecodeFile(RepoDataFile, &conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
}
