package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建两个无缓冲字符串类型通道 c1 和 c2。
	c1 := make(chan string)
	c2 := make(chan string)

	// 开启第一个 goroutine，在 1 秒后向 c1 通道中发送消息 "one"。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	// 开启第二个 goroutine，在 2 秒后向 c2 通道中发送消息 "two"。
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 循环两次，分别从 c1 和 c2 通道中接收消息并输出。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
