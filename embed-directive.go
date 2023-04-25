package main

import (
	"embed"
)

// 声明一个嵌入式字符串变量，用于嵌入 folder/single_file.txt 文件的内容
// 使用 go:embed 指令将文件嵌入到程序中
var fileString string

// 声明一个嵌入式字节数组变量，用于嵌入 folder/single_file.txt 文件的内容
// 使用 go:embed 指令将文件嵌入到程序中
var fileByte []byte

// 声明一个嵌入式文件系统变量，用于嵌入 folder 目录及其子目录中的所有 *.hash 文件
// 使用 go:embed 指令将文件夹嵌入到程序中
var folder embed.FS

func main() {

	// 打印嵌入式字符串变量 fileString 中的内容
	print(fileString)

	// 将嵌入式字节数组变量 fileByte 转换为字符串，并打印其内容
	print(string(fileByte))

	// 从嵌入式文件系统 folder 中读取并打印 folder/file1.hash 文件的内容
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	// 从嵌入式文件系统 folder 中读取并打印 folder/file2.hash 文件的内容
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
