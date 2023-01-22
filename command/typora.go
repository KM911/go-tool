package command

import "github.com/urfave/cli/v2"

var Typora = &cli.Command{
	Name:    "typora",
	Aliases: []string{"t"},
	Usage:   "open a soft",
	Action:  TyporaAction,
}

func TyporaAction(c *cli.Context) error {
	println("typora")
	return nil
}
