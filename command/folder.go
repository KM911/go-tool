package command

import (
	"github.com/urfave/cli/v2"
	"gt/config"
	"lib"
)

//
//
//gt folder a
//gt f a

var Folder = &cli.Command{
	Name:    "folder",
	Aliases: []string{"f"},
	Usage:   "open a folder",
	Action:  Gf,
}

func Gf(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		lib.Run("explorer.exe .")
		return nil
	}
	if len(c.Args().Slice()) == 1 {
		lib.Run("explorer.exe " + config.WorkFolderDict[c.Args().First()])
		return nil
	} else {
		println("too many argvs ")
		return nil
	}
	return nil
}