package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args //os.Args: 返回命令行参数的切片。
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
