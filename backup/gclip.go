package main

import (
	"lib"
	"os"
	"path/filepath"
)

func main() {
	pathName := os.Args[1]
	pathName = filepath.ToSlash(pathName)
	lib.Run("echo " + pathName + " | clip")
}

// D:\0.SOFT\A_script\bin\webp.exe  %1
