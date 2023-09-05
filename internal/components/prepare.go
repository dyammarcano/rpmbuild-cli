package components

import (
	"encoding/json"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/git"
	"log"
)

func Prepare(rootPath string) error {
	changelogs, err := git.Changelog(rootPath)
	if err != nil {
		log.Print(err)
	}

	jsonData, err := json.Marshal(changelogs)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}
