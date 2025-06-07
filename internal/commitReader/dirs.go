package commitReader

import (
	"os"
)

func CheckIsGitRepository() (bool, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return false, err
	}

	var gitInDir bool

	for _, file := range files {
		if file.Name() == ".git" {
			gitInDir = true
		}
	}

	return gitInDir, nil
}
