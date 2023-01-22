package main

import (
	"os"
	"os/exec"
)

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
func main() {
	dir, _ := os.Getwd()
	println("创建的项目名称为：", os.Args[1])
	println("创建的项目路径为：", dir+"\\", os.Args[1])
	projecPath := dir + "\\" + os.Args[1]
	Run("xcopy D:\\0.SOFT\\A_script\\vueinit " + projecPath + "\\ /sy")
	println("创建项目成功！")
}
