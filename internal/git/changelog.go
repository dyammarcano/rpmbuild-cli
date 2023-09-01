package git

import (
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Changelog(repoPath string) ([]structures.Changelog, error) {
	if err := os.Chdir(repoPath); err != nil {
		return nil, err
	}

	if _, err := os.Stat(filepath.Join(repoPath, ".git")); err == nil {
		output, err := exec.Command("git", "log", "--pretty=format:%h|%an|%ae|%ad|%s").Output()
		if err != nil {
			return nil, err
		}

		var commits []structures.Changelog

		for _, commitLine := range strings.Split(string(output), "\n") {
			commitDetails := strings.Split(commitLine, "|")
			if len(commitDetails) == 5 {
				commit := structures.Changelog{
					Hash:    commitDetails[0],
					Author:  commitDetails[1],
					Email:   commitDetails[2],
					Date:    commitDetails[3],
					Message: commitDetails[4],
				}
				commits = append(commits, commit)
			}
		}

		return commits, nil
	}

	fmt.Println("Not a git repository")
	return nil, nil
}
