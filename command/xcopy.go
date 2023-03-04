package command

import (
	"github.com/urfave/cli/v2"
	"gt/config"
	"lib"
	"os"
)

var (
	Xcopy = &cli.Command{
		Name:    "xcopy",
		Aliases: []string{"x"}, // conflict with install
		Usage:   "init a project by xcopy",
		Action:  XcopyAction,
	}
)

// 检查文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 参数个数的校验 这里是一个通用的操作
func XcopyAction(c *cli.Context) error {
	// 检查是否是已有的项目
	lens := len(c.Args().Slice())
	// 是否是ProjectList中的项目
	if lens == 0 {
		println("need a argv")
	} else if lens == 1 {

		project := config.ProjectPath + c.Args().First()
		if ok, _ := PathExists(project); ok {
			lib.Run("xcopy " + project + " ." + " /E /Y /I")
			if c.Args().First() == "git" {
				lib.Run("git init")
			}
		} else {
			println("invalid project name")
		}
	} else if lens == 2 {
		project := config.ProjectPath + c.Args().First()
		if ok, _ := PathExists(project); ok {
			lib.Run("xcopy " + project + " ./" + c.Args().Get(2) + " /E /Y /I")
			if c.Args().First() == "git" {
				lib.Run("git init")
			}
		} else {
			println("invalid project name")
		}

	}
	return nil
}
