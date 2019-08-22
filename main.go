package main

import (
	"fmt"
	"log"
	"os"

	"turbo-branch/git"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Turbo Branch"
	app.Usage = "Navigate your git branches with ease."
	app.Action = func(c *cli.Context) error {
		branches := git.GetBranches()
		fmt.Println(branches)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
