package main

import "fmt"

func main() {
	messages := make(chan string) // 创建一个新的通道，类型为string

	go func() { messages <- "Hello World" }() // 发送一个新的值到通道中，使用 <- 语法

	msg := <-messages
	fmt.Println(msg)
}
