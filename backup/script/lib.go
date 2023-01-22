// 这里就是记录我们之前写的一些基础工具
package main
// 这个是对于cil使用的
func Cmd(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// 直接运行指令
func Run(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Run()
}

// 这里我们执行指令并且返回值 对于指令正确性要求较高的时候可以使用 还可以监考错误输出
func RrturnValue(command string)  {
	cmd := exec.Command("cmd", "/C", command)
	
}
// 获取文件名 不含后缀
func WithoutExt(fileName string) string {
	ext := filepath.Ext(fileName)
	baseName := filepath.Base(filepath.ToSlash(fileName))
	return baseName[:len(baseName)-len(ext)]
}



import (
)

func main() {
	 
}