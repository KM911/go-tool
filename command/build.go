package command

import (
	"github.com/urfave/cli/v2"
	"lib"
)

var Build = &cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "build go project and run with argv",
	Action:  BuildAction,
}

func BuildAction(c *cli.Context) error {
	// 判断是否存在main.exe文件

	lens := len(c.Args().Slice())
	if lens == 0 {
		println("		build need a argv ")
		return nil
	} else {
		if lib.IsExit(lib.BaseName(c.Args().First()) + ".exe") {
			lib.Run("rm " + lib.BaseName(c.Args().First()) + ".exe")
		}
		// 这里就不应该检测是否是go文件 你就不应该写错好吧
		source := ""
		// 将每一个文件进行格式化
		for i := 0; i < lens; i++ {
			file := lib.ToGoFile(c.Args().Get(i))
			println("参数", i, file)
			source = source + " " + file
			// 这里是利用各种工具对代码进行格式化 避免出现很低级的错误 带来的性能开销是可以接受的范围
			lib.Run("goimports -w " + file)
			lib.Run("gofmt -w " + file)

		}
		// 这里我是想做类似于 联合编译的功能
		lib.Run("go build " + source)

	}
	// 尝试在这里获取Flag的值
	if c.Bool("run") {
		lib.Run(lib.BaseName(c.Args().Get(0)))
		return nil
	}
	// 以一个参数运行
	if c.String("argv") != "" {
		lib.Run(lib.BaseName(c.Args().Get(0)) + " " + c.String("argv"))
		return nil
	}
	return nil
}
