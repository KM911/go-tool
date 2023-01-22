package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"lib"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	HOST    = "http://81.68.91.70/"
	API     = "tinypicgo/upload"
	IsLocal = false
)

func ProcessName() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	return path.Dir(fullFilename)
}

// 参数1: 图片的url
// 参数2: 图片的保存路径
func DownloadFile(url string) string {
	if strings.Contains(url, HOST) {
		println("图片已经上传")
		//panic("已经上传过了 不必二次上传")
		os.Exit(1)
	}
	// 如果是本地文件的话 我们直接返回
	if url[:4] != "http" {
		IsLocal = true
		return url
	}
	// 根据上一次的经验 我们还需要写一个就是更加具有普世性的获取图片类型/名字的方法
	fileName := path.Base(url)
	savePath := "backup"
	dir := ProcessName()
	dir = filepath.ToSlash(dir)
	filepath := dir + "/" + savePath + "/"
	res, err := http.Get(url)
	if err != nil {
		println("下载失败")
		return url
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filepath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	io.Copy(writer, reader)
	return filepath + fileName
}

func UrlPrase(url string) {
	slash := filepath.ToSlash(url)
	// 将第一个 / 变成 //
	slash = strings.Replace(slash, "///", "//", 1)
	fmt.Println(slash)
}

// 输入图片的原始路径 返回转化后的图片路径
func ImageToWebp(imagePath string) string {
	ext := filepath.Ext(imagePath)
	if ext == ".webp" || ext == ".gif" {
		return imagePath
	}

	newName := lib.BaseName(imagePath) + ".webp"
	dir := ProcessName()
	newPath := filepath.ToSlash(path.Join(dir, "webp", newName))
	println("newpath", newPath)
	lib.Run("webp " + imagePath + " " + newPath)
	return newPath
}
func UploadImage(imagePath string) {
	// 读取图片的数据
	imageData := ReadImageAsBase64(imagePath)
	// 将时间戳作为图片的名称
	ext := filepath.Ext(imagePath)
	imageName := strconv.FormatInt(time.Now().Unix(), 10) + ext
	image := make(map[string]string)
	image["name"] = imageName
	image["data"] = imageData
	bytesData, err := json.Marshal(image)
	resp, err := http.Post(HOST+API, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		println("上传失败")
		panic(err)
		return
	}
	defer resp.Body.Close()
	// 读取响应的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("网络请求失败")
		panic(err)
		return
	}
	UrlPrase(string(body))

}
func ReadImageAsBase64(imagePath string) string {
	// 读取图片的数据
	file, err := os.Open(imagePath)
	if err != nil {
		println("图片读取失败")
		panic(err)
		return ""
	}
	defer file.Close()
	// 获取图片的大小
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	// 创建一个缓冲区
	buffer := make([]byte, fileSize)
	// 读取图片的数据
	file.Read(buffer)
	// 返回图片的数据
	return base64.StdEncoding.EncodeToString(buffer)
}

func main() {
	// 后面的参数是图片的url 可能是多个
	for _, v := range os.Args[1:] {
		imagePath := DownloadFile(v)
		newPath := ImageToWebp(imagePath)
		UploadImage(newPath)




		
		if IsLocal {
			// 移除本地的图片
			if path.Ext(imagePath) != ".gif" || path.Ext(imagePath) != ".webp" {
				os.Remove(imagePath)
			}
		}
	}
}

