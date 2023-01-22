package main

import (
	"gt/command"
	"gt/flag"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func ShowAllArgvs(c *cli.Context) error {
	println("All argvs:")
	for _, v := range c.Args().Slice() {
		println("argv  +:", v)
	}
	return nil
}
func main() {
	app := &cli.App{
		Name:     "gt",
		Usage:    "gt is a CLI tool for go ",
		Commands: command.CommandList,
		Flags:    flag.FlagList,
		Action: func(c *cli.Context) error {
			if len(c.Args().Slice()) == 0 {
				println("no argv please input argv")
				cli.ShowAppHelp(c)
				return nil
			} else {
				println("invalid argv please input valid argv")
				ShowAllArgvs(c)
				cli.ShowAppHelp(c)
				return nil
			}
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
