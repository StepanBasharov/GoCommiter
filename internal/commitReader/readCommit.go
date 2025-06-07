package commitReader

import (
	"errors"
	"os/exec"
	"strings"
)

type Commit struct {
	ModifyFiles     []string
	ModifyFilesData []string
}

func NewCommit() (*Commit, error) {
	isGit, err := CheckIsGitRepository()
	if err != nil {
		return nil, err
	}

	if isGit {
		return &Commit{}, nil
	} else {
		return nil, errors.New("not a git repository")
	}
}

func (c *Commit) ReadCommit() error {
	if err := c.getModifyFiles(); err != nil {
		return err
	}

	if err := c.GetModifyFilesChanges(); err != nil {
		return err
	}

	return nil
}

func (c *Commit) getModifyFiles() error {
	cmd := exec.Command("git", "diff", "--name-only", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}

	files := strings.Split(string(out), "\n")

	for idx := range files {
		files[idx] = strings.TrimSpace(files[idx])
	}

	c.ModifyFiles = files
	return nil

}

func (c *Commit) GetModifyFilesChanges() error {
	filesChanges := make([]string, 0)

	for _, file := range c.ModifyFilesData {
		cmd := exec.Command("git", "diff", file)
		out, err := cmd.Output()
		if err != nil {
			return nil
		}

		filesChanges = append(filesChanges, string(out))

	}

	return nil

}
