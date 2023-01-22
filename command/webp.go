package command

import "github.com/urfave/cli/v2"

var Webp = &cli.Command{
	Name:    "webp",
	Aliases: []string{"w"},
	Usage:   "convert a image to webp",
	Action:  WebpAction,
}

func WebpAction(c *cli.Context) error {
	println("webp")
	return nil
}
