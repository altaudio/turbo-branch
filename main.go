package main

import (
	"fmt"
	"log"
	"os"

	"turbo-branch/git"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Turbo Branch"
	app.Usage = "Navigate your git branches with ease."
	app.Action = func(c *cli.Context) error {
		branches := git.GetBranches()

		prompt := promptui.Select{
			Label: "Select a branch",
			Items: branches,
		}

		_, branch, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(branch)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
