package command

import (
	"github.com/urfave/cli/v2"
	"gt/config"
	"lib"
)

var Open = &cli.Command{
	Name:    "open",
	Aliases: []string{"o"},
	Usage:   "open a soft",
	Action:  Gm,
}

func Start(soft string) {
	// 我的简写是通过 字典的键和值同时进行匹配的
	for key, value := range config.ShortCutDict {
		if key == soft || value == soft {
			lib.Run("start " + config.ShortCutPath + value)
			return
		}
	}
	println("invalid argv")
}
func Gm(c *cli.Context) error {
	lens := len(c.Args().Slice())
	// 无参数
	if lens == 0 {
		println("open need a argv")
		return nil
	} else {
		for i := 0; i < lens; i++ {
			// 这里开始对我们的参数列表进行遍历
			Start(c.Args().Get(i))
		}

	}
	return nil
}
