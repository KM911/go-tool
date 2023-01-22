package command

import (
	"github.com/urfave/cli/v2"
	"gt/config"
	"lib"
)

// 作用是优化我们安装包的体验 可以直接使用名字安装包 而不是全部的路径
// 使用方式
// gt install gin == go get -u github.com/gin-gonic/ginq
// gt i GIN
// gt i gorm
// gt i cli

var Install = &cli.Command{
	Name:    "install",
	Aliases: []string{"i"},
	Usage:   "install a go package by short name instead of full url",
	Action:  InstallAction,
}

func InstallAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("no argv run go mod tidy")
		lib.Run("go mod tidy")
		return nil
	}
	for _, value := range c.Args().Slice() {
		lib.Run("go get -u " + config.InstallPackagePathDict[value])
	}
	return nil
}
