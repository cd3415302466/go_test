package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个缓冲区大小为 1 的字符串类型的 channel
	c1 := make(chan string, 1)

	// 启动一个 goroutine，在 2 秒后向 channel 中发送一个字符串
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
	}()

	// 使用 select 语句从 channel 中读取数据，如果 1 秒内没有数据可以读取，则超时
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 创建另一个缓冲区大小为 1 的字符串类型的 channel
	c2 := make(chan string, 1)

	// 启动一个 goroutine，在 2 秒后向 channel 中发送一个字符串
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "result 2"
	}()

	// 使用 select 语句从 channel 中读取数据，如果 3 秒内没有数据可以读取，则超时
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// Go语言并发编程
// channel 与 goroutine 的结合是 Go 语言的一大特色，它让 Go 语言的并发编程变得非常简单。
