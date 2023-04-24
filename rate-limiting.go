package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个缓冲区大小为5的channel
	requests := make(chan int, 5)
	// 向channel中发送5个请求
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	// 关闭channel，表示已经发送完毕
	close(requests)

	// 创建一个每200毫秒发送一个信号的ticker
	limiter := time.Tick(200 * time.Millisecond)

	// 从channel中读取请求，并在每次接收到limiter的一个信号时进行处理
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 创建一个缓冲区大小为3的channel
	burstyLimiter := make(chan time.Time, 3)

	// 用当前时间初始化这个channel
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 启动一个协程，每200毫秒发送一个信号到burstyLimiter中，保持channel中有3个元素
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 创建一个缓冲区大小为5的channel，向其中发送5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	// 关闭channel，表示已经发送完毕
	close(burstyRequests)

	// 从channel中读取请求，并在每次接收到burstyLimiter的一个信号时进行处理
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
