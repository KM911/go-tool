package command

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/urfave/cli/v2"
	"lib"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var Snap = &cli.Command{
	Name:    "snap",
	Aliases: []string{"s"},
	Usage:   "save the snap of video",
	Action:  SnapAction,
}

/*
提供视频的绝对路径 将其第一帧作为视频封面. 生成路径是当前的 就是有点搞笑了
*/
func GetSnapshot(videoPath string, __ISABS__ bool) string {
	buf := bytes.NewBuffer(nil)
	videoPath = filepath.ToSlash(videoPath)
	ffmpeg.Input(videoPath, ffmpeg.KwArgs{
		"loglevel": "quiet",
	}).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", int(1))}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	coverdata, err := imaging.Decode(buf)
	covername := GetVideoName(videoPath, __ISABS__)
	println(covername)
	err = imaging.Save(coverdata, covername)
	if err != nil {
		println("failed to save image:", err)
	}
	return covername[2:]
}

/*
生成封面的路径
*/
func GetVideoName(videoPath string, __ISABS__ bool) string {

	//return path.Join("static", "image", names+".jpg")
	if __ISABS__ {
		pathslice := strings.Split(videoPath, ".")
		pathslice = pathslice[:len(pathslice)-1]
		name := strings.Join(pathslice, ".") + ".jpg"
		return name
	} else {
		names := lib.BaseName(filepath.ToSlash(videoPath))
		names = strings.Split(names, ".")[0]
		return path.Join(lib.CmdPath(), names+".jpg")
	}
}
func SnapAction(c *cli.Context) error {
	lens := len(c.Args().Slice())
	// 无参数
	if lens == 0 {
		println("open need a argv")
		return nil
	} else if lens == 1 {
		// 只有视频文件路径
		videopath, _ := filepath.Abs(c.Args().First())
		videopath = filepath.ToSlash(videopath)
		println("abs path", videopath)
		if filepath.IsAbs(videopath) {
			GetSnapshot(videopath, true)
		} else {
			videopath = path.Join(lib.CmdPath(), videopath)
			GetSnapshot(videopath, false)
		}
	} else {
		println("to many args")
	}
	return nil
}
