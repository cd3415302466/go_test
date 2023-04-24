package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// readOp 结构体代表一个读操作
type readOp struct {
	key  int      // 要读取的键
	resp chan int // 发送读取响应的通道
}

// writeOp 结构体代表一个写操作
type writeOp struct {
	key  int       // 要写入的键
	val  int       // 要写入键的值
	resp chan bool // 发送写入响应的通道
}

func main() {

	var readOps uint64           // 计数器：读操作的数量
	var writeOps uint64          // 计数器：写操作的数量
	reads := make(chan readOp)   // 存储读请求的通道
	writes := make(chan writeOp) // 存储写请求的通道

	// 处理读和写请求的 goroutine
	// state 是一个用于存储系统当前状态的 map
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // 发送与键对应的值
			case write := <-writes:
				state[write.key] = write.val // 将值写入键
				write.resp <- true           // 发送成功响应
			}
		}
	}()

	// 启动多个 goroutine 来生成读请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),   // 在 0-4 之间生成一个随机键
					resp: make(chan int)} // 创建用于接收读取响应的通道
				reads <- read                 // 发送读请求
				<-read.resp                   // 等待来自 readOp 处理程序的响应
				atomic.AddUint64(&readOps, 1) // 原子性地增加 readOps 计数器
				time.Sleep(time.Millisecond)  // 引入延迟以模拟网络延迟
			}
		}()
	}

	// 启动多个 goroutine 来生成写请求
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),    // 在 0-4 之间生成一个随机键
					val:  rand.Intn(100),  // 在 0-99 之间生成一个随机值
					resp: make(chan bool)} // 创建用于接收写入响应的通道
				writes <- write                // 发送写请求
				<-write.resp                   // 等待来自 writeOp 处理程序的响应
				atomic.AddUint64(&writeOps, 1) // 原子性地增加 writeOps 计数器
				time.Sleep(time.Millisecond)   // 引入延迟以模拟网络延迟
			}
		}()
	}

	time.Sleep(time.Second) // 让系统运行一秒钟

	// 打印读和写计数器的最终值
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
