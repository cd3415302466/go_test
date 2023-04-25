package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")                 //os.Setenv(): 设置环境变量的值。
	fmt.Println("FOO:", os.Getenv("FOO")) //os.Getenv(): 获取环境变量的值。
	fmt.Println("BAR:", os.Getenv("BAR")) //os.Getenv(): 获取环境变量的值。

	fmt.Println()
	for _, e := range os.Environ() { //os.Environ(): 返回环境变量的切片。
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
