package git

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func getGitFilePath(root string) (string, error) {
	directories := strings.Split(root, "/")

	for len(directories) > 2 {
		gitFilePath := strings.Join(directories, "/") + "/.git"

		fileExists, err := exists(gitFilePath)
		if err != nil {
			return "", err
		}

		if fileExists {
			return gitFilePath, nil
		}

		directories = directories[:len(directories)-1]
	}

	return "", errors.New("getGitFilePath: Could not find a .git directory.")
}

func GetBranches() []string {
	callPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	gitFilePath, err := getGitFilePath(callPath)
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	b.WriteString(gitFilePath)
	b.WriteString("/refs/heads/")
	headsPath := b.String()

	var branches []string

	filepath.Walk(headsPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		branch := strings.Replace(path, headsPath, "", 1)
		branches = append(branches, branch)

		return err
	})

	return branches
}
