package main

import (
	"os"
	"os/exec"

	"path/filepath"
)

// 这个其实可以作为就是一个基础的工具
func ClipBoard(str string) {
	cmd := exec.Command("cmd", "/c", "echo "+str+" | clip")
	cmd.Run()
}
func main() {
	pathNaem := os.Args[1]
	pathNaem = filepath.ToSlash(pathNaem)
	ClipBoard(pathNaem)
}

// D:\0.SOFT\A_script\bin\webp.exe  %1
