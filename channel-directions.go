package main

import "fmt"

// ping函数接受一个只能发送数据的字符串类型通道和一个消息参数。
func ping(pings chan<- string, msg string) {
	// 将消息发送到pings通道中。
	pings <- msg
}

// pong函数接受一个只能接收数据的字符串类型通道和一个只能发送数据的字符串类型通道。
func pong(pings <-chan string, pongs chan<- string) {
	// 从pings通道中接收消息。
	msg := <-pings
	// 将消息发送到pongs通道中。
	pongs <- msg
}

func main() {
	// 创建一个缓冲区大小为1的字符串类型通道pings和pongs。
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// 向pings通道发送一条消息。
	ping(pings, "passed message")
	// 调用pong函数，将pings通道作为参数传递给它，并将返回结果存储在pongs通道中。
	pong(pings, pongs)
	// 从pongs通道中接收并打印消息。
	fmt.Println(<-pongs)
}
