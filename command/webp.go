package command

import (
	"github.com/urfave/cli/v2"
	"lib"
)

var Webp = &cli.Command{
	Name:    "webp",
	Aliases: []string{"w"},
	Usage:   "convert a image to webp",
	Action:  WebpAction,
}

// 将jpg png jpeg 转化为webp

// 这里我像写一种新的图片演示 就是avfi 这个是更加重要的内容
func WebpAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("no image path")
		return nil
	} else {
		imagepath := lib.AbsPath(c.Args().Slice()[0])
		if lib.IsExit(imagepath) {
			lib.Run("webp " + imagepath)
		} else {
			println("image is not exit")
			println("please check your image path")
		}
		return nil
	}

}

// squoosh-cli  --avif  auto -d test  2.jpg 将其 avif 这个是js文件 很厉害
// webp 可以将图片变成 webp
// ffmpeg 提取视频
