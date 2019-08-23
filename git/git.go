package git

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetBranches() []string {
	callPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	b.WriteString(callPath)
	b.WriteString("/.git/refs/heads/")
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
