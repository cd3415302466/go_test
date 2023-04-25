package main

import (
	"bufio"    // 提供带缓冲的 I/O 功能
	"fmt"      // 提供格式化文本输出功能
	"net/http" // 提供 HTTP 客户端和服务器实现
)

func main() {

	// 发送 GET 请求并获取响应
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err) // 如果出现错误，则直接停止程序并打印错误信息
	}
	defer resp.Body.Close() // 关闭响应体以释放资源

	fmt.Println("Response status:", resp.Status) // 打印响应状态码

	// 使用 Scanner 逐行读取响应主体中的前五行文本
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	// 如果在扫描器上调用了错误方法，则 panic
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
