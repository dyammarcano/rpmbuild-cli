package git

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/pkg/errors"
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
		output, err := exec.Command("git", "log", "--pretty=format:%h|%an|%ae|%ad|%B<#>").Output()
		if err != nil {
			return nil, err
		}

		var commits []structures.Changelog

		for _, commitLine := range strings.Split(string(output), "<#>") {
			commitDetails := strings.Split(commitLine, "|")
			if len(commitDetails) == 5 {
				parts := strings.SplitN(commitDetails[4], "\n\n", 2)

				commit := structures.Changelog{
					Hash:   strings.TrimSpace(commitDetails[0]),
					Author: strings.TrimSpace(commitDetails[1]),
					Email:  strings.TrimSpace(commitDetails[2]),
					Date:   strings.TrimSpace(commitDetails[3]),
					Title:  strings.TrimSpace(parts[0]),
				}

				if len(parts) == 2 {
					commit.Message = strings.TrimSpace(parts[1])
				}

				commits = append(commits, commit)
			}
		}

		return commits, nil
	}

	return nil, errors.Errorf("not a git repository")
}
