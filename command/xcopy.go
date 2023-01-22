package command

import (
	"github.com/urfave/cli/v2"
	"gt/config"
	"lib"
)

var (
	Xcopy = &cli.Command{
		Name: "xcopy",
		Aliases: []string{"x"}, // conflict with install
		Usage:  "init a project by xcopy",
		Action: XcopyAction,
	}
)

func XcopyAction(c *cli.Context) error {
	// 检查是否是已有的项目
	for _, v := range config.ProjectList {
		if c.Args().First() == v {
			project := config.ProjectPath + v
			println("project:", project)
			
			lib.Run("xcopy " + project + " ." + " /E /Y /I")
			if c.Args().First() == "git" {
				lib.Run("git init")
			}
			return nil
		}
	}
	println("invalid project name")
	return nil

}
