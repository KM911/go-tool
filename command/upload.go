package command

import (
	"github.com/urfave/cli/v2"
	"lib"
)

// TODO Upload
var Upload = &cli.Command{
	Name:    "Upload",
	Aliases: []string{"u"},
	Usage:   "upload image to the tinypicgo",
	Action:  UploadAction,
}

func UploadAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("no argv run go mod tidy")
		lib.Run("go mod tidy")
		return nil
	}
	for _, value := range c.Args().Slice() {
		lib.Run("gi " + value)
	}
	return nil
}

/*
将图片上传到云端
*/
func UploadFile(__FILE__PATH__ string) {
	if lib.IsExit(__FILE__PATH__) {
		println("file is exit")
	} else {
		println("file is not exit")
	}
}
