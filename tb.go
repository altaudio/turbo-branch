package main

import (
	"log"
	"os"
	"os/exec"

	"turbo-branch/git"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// TODO: Find .git directory recursively
// TODO: Only getting the first part of the branch name :(

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

		cmd := exec.Command("git", "checkout", branch)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
