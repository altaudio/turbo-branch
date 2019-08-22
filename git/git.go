package git

import (
	"io/ioutil"
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
	b.WriteString("/.git/refs/heads")

	files, err := ioutil.ReadDir(b.String())
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames
}
