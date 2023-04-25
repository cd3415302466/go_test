package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string") //flag.String(): 定义一个字符串类型的命令行参数，返回一个指向该参数值的指针。
	numbPtr := flag.Int("numb", 42, "an int")         //flag.Int(): 定义一个整型类型的命令行参数，返回一个指向该参数值的指针。
	forkPtr := flag.Bool("fork", false, "a bool")     //flag.Bool(): 定义一个布尔类型的命令行参数，返回一个指向该参数值的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") //flag.StringVar(): 定义一个字符串类型的命令行参数，返回一个指向该参数值的指针。
	flag.Parse()                                         //flag.Parse(): 解析命令行参数。

	fmt.Println("word:", *wordPtr)    //*wordPtr: 获取指针指向的值。
	fmt.Println("numb:", *numbPtr)    //*numbPtr: 获取指针指向的值。
	fmt.Println("fork:", *forkPtr)    //*forkPtr: 获取指针指向的值。
	fmt.Println("svar:", svar)        //svar: 获取指针指向的值。
	fmt.Println("tail:", flag.Args()) //flag.Args(): 返回命令行参数后的其他参数。
}
