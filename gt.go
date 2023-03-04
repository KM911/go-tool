package main

import (
	"github.com/urfave/cli/v2"
	"gt/command"
	"gt/flag"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "gt",
		Usage:    "gt is a CLI tool for go ",
		Commands: command.CommandList,
		Flags:    flag.FlagList,
		Action: func(c *cli.Context) error {
			println("need a valid command")
			cli.ShowAppHelp(c)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
