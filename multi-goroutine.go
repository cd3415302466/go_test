package main

import (
	"fmt"
	"sync"
)

// printNumbers 函数接受一个指向 sync.WaitGroup 值的指针 wg 和一个整型 id。
func printNumbers(wg *sync.WaitGroup, id int) {
	defer wg.Done() // 在函数退出时通知 WaitGroup 完成任务。

	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup // 创建一个 WaitGroup 值。

	for i := 1; i <= 3; i++ { // 循环创建 3 个 goroutine。
		wg.Add(1)               // 向 WaitGroup 中添加一个等待任务。
		go printNumbers(&wg, i) // 开启一个 goroutine 调用 printNumbers 函数，将 WaitGroup 的地址和当前循环变量 i 作为参数传递给它。
	}

	// 等待所有 goroutine 执行完毕。
	wg.Wait()

	fmt.Println("All goroutines complete.")
}
