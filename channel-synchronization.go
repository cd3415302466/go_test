package main

import (
	"fmt"
	"time"
)

// worker 函数接受一个只能发送 bool 类型数据的通道 done，表示任务是否完成。
func worker(done chan bool) {
	fmt.Print("working...") // 输出 "working..."
	time.Sleep(time.Second) // 模拟工作
	fmt.Println("done")     // 输出 "done"

	done <- true // 向通道 done 发送任务已完成的消息。
}

func main() {
	done := make(chan bool, 1) // 创建缓冲区大小为1的 bool 类型通道 done。
	go worker(done)            // 开启一个协程调用 worker 函数，将通道 done 作为参数传递给它。
	<-done                     // 阻塞主程序，等待从通道 done 中接收到任务完成的消息。
}
