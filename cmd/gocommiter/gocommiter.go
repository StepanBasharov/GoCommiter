package main

import (
	"fmt"
	"github.com/urfave/cli"
	"gocommiter/internal/commitMaker"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "GoCommiter",
		Usage: "commit changes to git",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "commit",
				Usage:    "Commit description",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			commit, err := commitMaker.NewCommitMaker()
			if err != nil {
				return err
			}

			commitDescription := c.String("commit")

			if err := commit.MakeCommit(commitDescription); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("❌ Ошибка:", err)
	}
}
