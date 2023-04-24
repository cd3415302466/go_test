package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建定时器 timer1，2 秒后触发
	timer1 := time.NewTimer(2 * time.Second)

	// 等待 timer1 的信号（<-timer1.C），当信号到达时打印信息
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 创建定时器 timer2，1 秒后触发，并开启一个 goroutine 监听信号
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// 停止定时器 timer2，如果成功则打印信息
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 暂停执行 2 秒钟以保证所有 goroutine 完成执行
	time.Sleep(2 * time.Second)
}
