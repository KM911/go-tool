package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

func DirName() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	fullFilename = strings.Replace(fullFilename, "bin", "", 1)
	return path.Dir(filepath.ToSlash(fullFilename))
}

var (
	kernel32DLL                  = syscall.MustLoadDLL("User32.dll")
	procEnumWindows              = kernel32DLL.MustFindProc("EnumWindows")
	procGetWindowText            = kernel32DLL.MustFindProc("GetWindowTextW")
	procGetWindowTextLength      = kernel32DLL.MustFindProc("GetWindowTextLengthW")
	procGetWindowLong            = kernel32DLL.MustFindProc("GetWindowLongW")
	procIsWindowVisible          = kernel32DLL.MustFindProc("IsWindowVisible")
	procGetWindowThreadProcessId = kernel32DLL.MustFindProc("GetWindowThreadProcessId")
	k2                           = syscall.MustLoadDLL("kernel32.dll")
	psapi                        = syscall.MustLoadDLL("psapi.dll")
	procGetModuleFileNameEx      = psapi.MustFindProc("GetModuleFileNameExW")

	procOpenProcess           = k2.MustFindProc("OpenProcess")
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_VM_READ           = 0x0010

	procQueryFullProcessImageName = k2.MustFindProc("QueryFullProcessImageNameW")

	Dir = DirName()
)

func Run(command string) {
	// cmd := syscall.StringToUTF16Ptr(command)
	// syscall.Syscall(procGetModuleFileNameEx.Addr(), 4, uintptr(0), uintptr(0), uintptr(unsafe.Pointer(cmd)), uintptr(0))
	cmd := exec.Command("cmd", "/C", command)
	cmd.Run()
}
func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0)
	return &chars[0]
}

// 回调函数，用于EnumWindows中的回调函数，第一个参数是hWnd，第二个是自定义穿的参数
func AddElementFunc(hWnd syscall.Handle, hWndList *[]syscall.Handle) uintptr {
	*hWndList = append(*hWndList, hWnd)
	return 1
}

// 判读全部窗口是否可见
func IsWindowVisible(hWnd syscall.Handle) bool {
	r1, _, err := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hWnd), 0, 0)
	if err != 0 {
		fmt.Println(err)
	}
	return r1 != 0
}
func SaveWorkSpace() {
	// 创建一个数组，用于存储所有窗口句柄
	var PidList [20]int
	count := 0
	var hWndList []syscall.Handle
	hL := &hWndList
	syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(syscall.NewCallback(AddElementFunc)), uintptr(unsafe.Pointer(hL)), 0)

	// 遍历所有窗口 看是否可见
	for _, hWnd := range hWndList {
		if IsWindowVisible(hWnd) {
			// 获取窗口标题长度
			r1, _, err := syscall.Syscall(procGetWindowTextLength.Addr(), 1, uintptr(hWnd), 0, 0)
			if err != 0 {
				fmt.Println(err)
			}
			// 获取窗口标题
			var buffer [1024]uint16
			syscall.Syscall(procGetWindowText.Addr(), 3, uintptr(hWnd), uintptr(unsafe.Pointer(&buffer[0])), r1+1)
			title := syscall.UTF16ToString(buffer[:])
			if title == "" || title == "Default IME" || title == "MSCTFIME UI" || title == "Program Manager" || title == "DWM Notification Window" || title == "DDE Server Window" || title == "Settings" || title == "Microsoft Text Input Application" {
				continue
			}

			// 通过窗口句柄获取进程ID
			var processID uint32
			syscall.Syscall(procGetWindowThreadProcessId.Addr(), 2, uintptr(hWnd), uintptr(unsafe.Pointer(&processID)), 0)
			// 这里其实我可以利用就是cmd获取返回的信息就好了

			PidList[count] = int(processID)
			count++

		}
	}
	argv := ""
	// 遍历 pidlist
	for i := 0; i < count; i++ {
		argv += " " + fmt.Sprintf("%d", PidList[i])
	}
	Run("SaveWorkSpace.bat" + argv)

}

func ReSrotageWorkSpace() {
	// 解析我们的data.txt 文件 利用gm 进行启动就好了
	// 读取data.txt 文件  利用slice 获取我们的全部应用进程
	// 通过gm 进行启动
	// 利用bufio 读取文件
	file, err := os.Open(path.Join(Dir, "script", "data.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		softList := strings.Split(scanner.Text(), ",")
		for index, soft := range softList {
			if index == len(softList)-1 {
				break
			}
			println("gm " + strings.Replace(strings.Replace(soft, "\"", "", -1), ".exe", "", 1))
			Run("gm " + strings.Replace(strings.Replace(soft, "\"", "", -1), ".exe", "", 1))
			//println()
		}

	}

}

func main() {
	if os.Args[1] == "save" || os.Args[1] == "s" {
		SaveWorkSpace()
		Run("shutdown -s -t 10")
		println("save work space and shutdown")
	}
	if os.Args[1] == "restore" || os.Args[1] == "r" {
		ReSrotageWorkSpace()
	}
	if os.Args[1] == "reset" || os.Args[1] == "rs" {
		SaveWorkSpace()
		Run("shutdown -r -t 10")

	}
	if os.Args[1] == "help" || os.Args[1] == "h" {
		println("input with save or s to save work space and shutdown")
		println("input with restore or r to restore work space")
	}
}
