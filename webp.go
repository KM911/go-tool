package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Cwebp(input string, output string) {
	cmd := exec.Command("cmd", "/c", "D:\\1.CS\\git_hub\\script\\node_modules\\webp-converter\\bin\\libwebp_win64\\bin\\cwebp.exe -q 80 "+input+" -o "+output)
	cmd.Run()
}
func main() {
	if len(os.Args) == 1 {
		println("请输入要转换的图片")
		return
	}
	filePath := os.Args[1]
	// 如果是绝对路径
	if filePath[1] == ':' {
		println("绝对路径")
	} else {
		// 如果是相对路径
		filePath, _ = filepath.Abs(filePath)
		println("相对路径")
	}

	fileName := filepath.Base(filePath)

	// 判端是否输入了文件名
	if os.Args == nil || len(os.Args) < 2 {
		println("请输入要转换的图片")
		return
	}
	_, err := os.Stat(filePath)
	if err != nil {
		println("文件不存在")
		return
	}
	// 如果文件就是webp格式，就不需要转换了
	if fileName[len(fileName)-4:] == "webp" || fileName[len(fileName)-4:] == ".gif" {
		println("无需转换")
		return
	}
	// 转换为webp格式

	fileNameWithoutExt := fileName[:len(fileName)-4]
	webpFilePath := fileNameWithoutExt + ".webp"

	if len(os.Args) > 2 {
		if os.Args[2] == "d" {
			os.Remove(filePath)
			println("删除原文件")
		}

		// 转换为webp格式

		if os.Args[2][1] == ':' {
			webpFilePath = filepath.ToSlash(os.Args[2])
			// 这里就是将文件输出到指定的目录
		}
	}
	Cwebp(filePath, webpFilePath)
}
