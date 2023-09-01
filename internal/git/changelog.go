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
		output, err := exec.Command("git", "log", "--pretty=format:%h|%an|%ae|%ad|%B").Output()
		if err != nil {
			return nil, err
		}

		var commits []structures.Changelog

		for _, commitLine := range strings.Split(string(output), "\n\n") {
			commitDetails := strings.Split(commitLine, "|")
			if len(commitDetails) == 5 {
				commit := structures.Changelog{
					Hash:    strings.TrimSpace(commitDetails[0]),
					Author:  strings.TrimSpace(commitDetails[1]),
					Email:   strings.TrimSpace(commitDetails[2]),
					Date:    strings.TrimSpace(commitDetails[3]),
					Message: strings.TrimSpace(strings.Join(commitDetails[4:], "\n")),
				}
				commits = append(commits, commit)
			}
		}

		return commits, nil
	}

	fmt.Println("Not a git repository")
	return nil, nil
}
