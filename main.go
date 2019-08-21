package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Turbo Branch"
	app.Usage = "Navigate your git branches with ease."
	app.Action = func(c *cli.Context) error {
		fmt.Println("some branches")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
