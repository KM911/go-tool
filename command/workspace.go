package command

import (
	"github.com/urfave/cli/v2"
	"lib"
)

var WorkSpace = &cli.Command{
	Name:    "WorkSpace",
	Aliases: []string{"w"},
	Usage:   "save your workspace(the running software) while shutdown, and open them when you need them",
	Action:  WorkSpaceAction,
}

func WorkSpaceAction(c *cli.Context) error {
	lens := len(c.Args().Slice())
	if lens == 0 {
		println("no argv")
		return nil

	} else if lens == 1 {
		if c.Args().First() == "save" || c.Args().First() == "s" {
			lib.Run("ws s")
		} else if c.Args().First() == "restore" || c.Args().First() == "r" {
			lib.Run("ws r")
		}
	} else {
		println("too many argv")
	}
	return nil
}
