package command

import (
	"github.com/urfave/cli/v2"
	"lib"
)

var Clip = &cli.Command{
	Name:    "Clip",
	Aliases: []string{"c"},
	Usage:   "clip a file absolute path",
	Action:  ClipAction,
}

func ClipAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("need a argv")
		return nil
	} else {
		lib.Run("echo " + lib.AbsPath(c.Args().First()) + " | clip")
	}
	return nil
}
