package flag

import (
	"github.com/urfave/cli/v2"
)

var RunAfterBuild = &cli.BoolFlag{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "run after build",
}

var RunWithArgv = &cli.StringFlag{
	Name:    "argv",
	Aliases: []string{"a"},
	Usage:   "run with argv",
}
