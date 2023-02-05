package command

import "github.com/urfave/cli/v2"

var WorkSpace = &cli.Command{
	Name:    "WorkSpace",
	Aliases: []string{"w"},
	Usage:   "save your workspace(the running software) while shutdown, and open them when you need them",
	Action:  WorkSpaceAction,
}

func WorkSpaceAction(c *cli.Context) error {
	println("typora")
	return nil
}
