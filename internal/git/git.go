package git

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/structures"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"strings"
)

type (
	Git struct {
	}

	Info struct {
		CommitHash  string
		ProjectName string
		LastTag     string
		Version     string
		RemoteRepo  string
		Changelog   []structures.Changelog
	}
)

// GetGitInfo retrieves Git information from a local repository.
func GetGitInfo(path string) (*Info, error) {
	// Open the Git repository at the specified path.
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	// Get the HEAD reference (current branch).
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	// Get the commit hash of the HEAD reference.
	commitHash := ref.Hash().String()

	// Get the project name from the repository's URL.
	remote, err := r.Remote("origin")
	if err != nil {
		return nil, err
	}
	projectName := extractProjectName(remote.Config().URLs[0])

	// Get the latest tag in the repository.
	tags, err := r.Tags()
	if err != nil {
		return nil, err
	}

	var lastTag string
	_ = tags.ForEach(func(tag *plumbing.Reference) error {
		lastTag = tag.Name().Short()
		return nil
	})

	changelog, err := Changelog(path)
	if err != nil {
		return nil, err
	}

	// Determine the version based on the last tag.
	version := lastTag

	gitInfo := &Info{
		CommitHash:  commitHash,
		ProjectName: projectName,
		LastTag:     lastTag,
		Version:     version,
		RemoteRepo:  strings.TrimSuffix(remote.Config().URLs[0], ".git"),
		Changelog:   changelog,
	}

	return gitInfo, nil
}

// extractProjectName extracts the project name from a Git repository URL.
func extractProjectName(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 2 {
		return strings.TrimSuffix(parts[len(parts)-1], ".git")
	}
	return url // Fallback to the full URL if extraction fails.
}

func NewGit(rootPath string) *Git {
	//// Clones the given repository in memory, creating the remote, the local
	//// branches and fetching the objects, exactly as:
	//Info("git clone https://github.com/go-git/go-billy")
	//
	//r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
	//	URL: "https://github.com/go-git/go-billy",
	//})
	//
	//CheckIfError(err)
	//
	//// Gets the HEAD history from HEAD, just like this command:
	//Info("git log")
	//
	//// ... retrieves the branch pointed by HEAD
	//ref, err := r.Head()
	//CheckIfError(err)
	//
	//// ... retrieves the commit history
	//cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	//CheckIfError(err)
	//
	//// ... just iterates over the commits, printing it
	//err = cIter.ForEach(func(c *object.Commit) error {
	//	fmt.Println(c)
	//	return nil
	//})
	//CheckIfError(err)
	return &Git{}
}

//func (g *Git) Changelog() {
//
//}
//
//func (g *Git) Branch() {
//
//}
//
//func (g *Git) Commit() {
//
//}
//
//func (g *Git) Tag() {
//
//}
//
//func (g *Git) Push() {
//
//}
//
//func (g *Git) Pull() {
//
//}
//
//func (g *Git) Merge() {
//
//}
//
//func (g *Git) Rebase() {
//
//}
//
//func (g *Git) Reset() {
//
//}
//
//func (g *Git) Status() {
//
//}
//
//func (g *Git) Diff() {
//
//}
//
//func (g *Git) Add() {
//
//}
//
//func (g *Git) Remove() {
//
//}
//
//func (g *Git) Stash() {
//
//}
//
//func (g *Git) Clean() {
//
//}
//
//func (g *Git) Init() {
//
//}
//
//func (g *Git) Clone() {
//
//}
//
//func (g *Git) Fetch() {
//
//}
//
//func (g *Git) MergeBase() {
//
//}
//
//func (g *Git) Archive() {
//
//}
//
//func (g *Git) Submodule() {
//
//}
//
//func (g *Git) Worktree() {
//
//}
//
//func (g *Git) Remote() {
//
//}
//
//func (g *Git) Config() {
//
//}
//
//func (g *Git) Describe() {
//
//}
//
//func (g *Git) Grep() {
//
//}
//
//func (g *Git) Log() {
//
//}
//
//func (g *Git) Notes() {
//
//}
//
//func (g *Git) Rm() {
//
//}
//
//func (g *Git) Show() {
//
//}
//
//func (g *Git) TagObject() {
//
//}
//
//func (g *Git) VerifyCommit() {
//
//}
//
//func (g *Git) VerifyTag() {
//
//}
//
//func (g *Git) WorktreeAdd() {
//
//}
//
//func CheckIfError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
