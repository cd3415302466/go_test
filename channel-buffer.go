package main

import "fmt"

func main() { // 通道缓冲
	messages := make(chan string, 2) // 创建一个缓冲通道，最多允许缓冲两个值

	messages <- "Hello World"
	messages <- "Bye World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
