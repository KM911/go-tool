package main

import (
	"fmt"
	"os"
	"os/exec"
)

func InitShortCutDict(shortCutDict map[string]string) {
	shortCutDict["tf"] = "tencentfile"
	shortCutDict["tx"] = "tencentfile"
	shortCutDict["a"] = "apifox"
	shortCutDict["b"] = "baidu"
	shortCutDict["bl"] = "blogs"
	shortCutDict["car"] = "Carnac"
	shortCutDict["bc"] = "bcut"
	shortCutDict["tu"] = "tuba"

	shortCutDict["g"] = "chrome"
	shortCutDict["ch"] = "chrome"

	shortCutDict["cl"] = "Clash"
	shortCutDict["cla"] = "Clash"
	shortCutDict["clash"] = "Clash"

	shortCutDict["cm"] = "ContextMenu"
	shortCutDict["dl"] = "deepl"
	shortCutDict["di"] = "dict"

	shortCutDict["doc"] = "docker"
	shortCutDict["do"] = "docker"

	shortCutDict["ep"] = "epic"
	shortCutDict["every"] = "everything"
	shortCutDict["er"] = "everything"

	shortCutDict["fm"] = "explorer"
	shortCutDict["ex"] = "explorer"
	shortCutDict["f"] = "explorer"

	shortCutDict["ftp"] = "filezilla"

	shortCutDict["test"] = "finallytest"
	shortCutDict["ft"] = "finallytest"

	shortCutDict["fox"] = "firefox"

	shortCutDict["github"] = "github"
	shortCutDict["hub"] = "git_hub"
	shortCutDict["gl"] = "goland"
	shortCutDict["go"] = "goland"
	shortCutDict["gm"] = "Grammarly"
	shortCutDict["id"] = "idea"
	shortCutDict["i"] = "idea"

	shortCutDict["md"] = "markdown"
	shortCutDict["mm"] = "mindmaster"
	shortCutDict["db"] = "mongodb"
	shortCutDict["m"] = "mongodb"

	shortCutDict["mo"] = "motrix"
	shortCutDict["mx"] = "motrix"

	shortCutDict["na"] = "navicat"
	shortCutDict["n"] = "navicat"
	shortCutDict["cat"] = "navicat"

	shortCutDict["nut"] = "nutstore"
	shortCutDict["ns"] = "nutstore"

	shortCutDict["vpn"] = "okztwo"
	shortCutDict["ok"] = "okztwo"

	shortCutDict["pcha"] = "PigchaProxy"
	shortCutDict["vps"] = "PigchaProxy"
	shortCutDict["pig"] = "PigchaProxy"

	shortCutDict["rb"] = "recyle"

	shortCutDict["st"] = "steam"
	shortCutDict["t"] = "tim"
	shortCutDict["to"] = "todo"

	shortCutDict["ty"] = "typora"
	shortCutDict["ub"] = "ubisoft"

	shortCutDict["ul"] = "ulearning"
	shortCutDict["u"] = "ulearning"

	shortCutDict["w"] = "wechat"
	shortCutDict["wall"] = "wallpaper"
	shortCutDict["wp"] = "wallpaper"
	shortCutDict["wt"] = "watt"
	shortCutDict["wat"] = "watt"

	shortCutDict["we"] = "wechat"
	shortCutDict["w"] = "wechat"
	shortCutDict["wx"] = "wechat"

	shortCutDict["meet"] = "wemeet"
	shortCutDict["wemeet"] = "wemeetapp"
	shortCutDict["wm"] = "wemeet"
	shortCutDict["fs"] = "feishu"

	// 文件夹
	shortCutDict["d"] = "D"
	shortCutDict["c"] = "C"
	shortCutDict["e"] = "E"
}

func RunCommand(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
func Run(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Run()
}

// 我们还需要一个就是将字符串全部转换为小写的函数
func toLower(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] + 32
		}
	}
	return string(b)
}
func help_info(shortCutDict map[string]string, folderDict map[string]string) {
	println("帮助信息 输入简称和全写都可以识别 例如：gm g 或者 gm chrome")
	println("打开软件的快捷键")
	count := 0
	for key, value := range shortCutDict {
		count++
		if count%5 == 0 {
			println()
		}
		fmt.Printf("%-10s", key+" - "+value+"\t")
	}
	println()
	println("打开文件夹的快捷键")
	for key, value := range folderDict {
		println(key, " - ", value)
	}
}
func main() {
	argv := toLower(os.Args[1])
	shortCutDict := make(map[string]string)
	shortCutPath := "D:\\0.SOFT\\A_shortcut\\"
	InitShortCutDict(shortCutDict)

	// 软件

	// 判断参数是否为空
	if len(os.Args) == 1 {
		println("参数不能为空！")
		return
	}
	// 特殊指令部分
	codeFolderDict := make(map[string]string)
	codeFolderDict["go"] = "D:\\1.CS\\Coding\\GO"
	codeFolderDict["c"] = "D:\\1.CS\\Coding\\C"
	codeFolderDict["cpp"] = "D:\\1.CS\\Coding\\C++"
	codeFolderDict["py"] = "D:\\1.CS\\Coding\\Python"
	codeFolderDict["js"] = "D:\\1.CS\\Coding\\Node"
	codeFolderDict["node"] = "D:\\1.CS\\Coding\\Node"
	codeFolderDict["java"] = "D:\\1.CS\\Coding\\Java"
	// 这个是用vscode 来打开我们的代码文件夹
	codeFolderDict["bin"] = "D:\\0.SOFT\\Code\\VsCode.exe --extensions-dir D:\\0.SOFT\\Code\\extensions  --user-data-dir D:\\0.SOFT\\Code\\data D:\\1.CS\\Coding\\GO\\go-tool"
	if argv == "code" || argv == "vscode" {
		if len(os.Args) == 2 {
			Run("start D:\\0.SOFT\\Code\\VsCode.exe --extensions-dir D:\\0.SOFT\\Code\\extensions  --user-data-dir D:\\0.SOFT\\Code\\data ")
			return
		}
		// 需要为我们打开一个文件夹 和 vscode的界面
		folder := os.Args[2]
		// 这里我们还需要维护一个就是代码文件夹的字典

		// 差不多是这个样子的
		for k, v := range codeFolderDict {
			if k == folder {
				if k != "bin" {
					Run("cf.vbs " + v)
				} else {
					Run("cf.vbs " + "D:\\1.CS\\Coding\\Go\\go-tool")
				}
				Run("start " + v)

				return
			}
		}
	}
	if argv == "help" {
		help_info(shortCutDict, codeFolderDict)
		return
	}

	// 执行快捷方式部分
	// 判断参数是否在字典中
	if _, ok := shortCutDict[argv]; !ok {
		// 遍历字典的value 看参数是否在字典的value中
		for _, v := range shortCutDict {
			if v == argv {
				Run("start " + shortCutPath + argv)
				return
			}
		}
		println("无效参数！输入help可以查看全部的参数")
	} else {
		// 打开快捷方式
		Run("start " + shortCutPath + shortCutDict[argv])
	}
}
