package commitMaker

import (
	"errors"
	"fmt"
	"gocommiter/internal/commitReader"
	"os/exec"
)

// Make commit

type Maker struct {
	commit *commitReader.Commit
}

func NewCommitMaker() (Maker, error) {
	commit, err := commitReader.NewCommit()
	if err != nil {
		return Maker{}, err
	}

	if commit.ModifyFiles == nil && &commit.ModifyFilesData == nil {
		return Maker{}, errors.New("nothing to commit")
	}

	return Maker{commit: commit}, nil
}

func (m Maker) MakeCommit(commitDescription string) error {
	if commitDescription == "" {
		commitDescription = m.createCommitDescription()
	}

	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", commitDescription)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "push")
	if err := cmd.Run(); err != nil {
		return err
	}

	m.printCommitInfo()

	return nil
}

func (m Maker) createCommitDescription() string {
	text := "The following files have been modified in this commit:\n\n"

	for _, file := range m.commit.ModifyFiles {
		text += file + "\n"
	}

	return text
}

func (m Maker) printCommitInfo() {
	changed := m.createCommitDescription()

	fmt.Println(changed + "\n\n")

	for _, fileChanges := range m.commit.ModifyFilesData {
		fmt.Println(fileChanges + "\n\n")
	}
}
