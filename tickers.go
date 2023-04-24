package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个每500毫秒执行一次的ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	// 创建一个信道用于通知goroutine停止运行
	done := make(chan bool)

	// 启动一个goroutine来处理ticker事件和停止信道
	go func() {
		for {
			select {
			// 如果收到了停止信号，goroutine退出
			case <-done:
				return
			// 处理ticker事件
			case t := <-ticker.C:
				fmt.Println("ticker at", t)
			}
		}
	}()

	// 让主goroutine等待2.1秒钟
	time.Sleep(2100 * time.Millisecond)
	// 停止ticker
	ticker.Stop()
	// 发送停止信号
	done <- true
	// 输出"Ticker stopped"
	fmt.Println("Ticker stopped")
}
