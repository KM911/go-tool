// 1 格式化代码  gofmt -w main.go
// 2 将代码进行编译 go build main.go

// 这里我们需要输入两次指令 我们进行简化一下 只需要输入文件名就可以了

// 这里还有一个痛点 就是说 我希望就是获取参数的app 需要先build 然后再运行 这里我们直接可以携带参数运行

package main

import (
	"os"
	"os/exec"
	"path"
)

func Run(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	// 获取文件名
	if len(os.Args) == 1 {
		println("请输入要编译的文件名")
		return
	}
	fileName := os.Args[1]
	ext := path.Ext(fileName)
	fileName = fileName[0 : len(fileName)-len(ext)]
	if len(os.Args) >= 2 {
		// 将文件名进行编译
		Run("gofmt -w " + fileName + ".go")
		Run("go build  -o bin/" + path.Base(fileName) + ".exe " + fileName + ".go")

	}
	if len(os.Args) == 3 {
		// 这里是最好还会有一个参数 用来运行时添加参数
		Run("bin/" + path.Base(fileName) + ".exe " + os.Args[2])
	}
}
