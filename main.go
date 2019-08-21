package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Turbo Branch"
	app.Usage = "Navigate your git branches with ease."
	app.Action = func(c *cli.Context) error {
		// You can find the git branches in .git/refs/heads
		callPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(callPath)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
