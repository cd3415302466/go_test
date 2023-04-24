package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 在goroutine中使用新变量i保存当前循环变量的值
		i := i

		go func() {
			defer wg.Done()

			// 调用worker函数并传递worker id参数
			worker(i)
		}()
	}

	// 等待所有goroutine执行完毕
	wg.Wait()
}

//如果不加 i := i，则在goroutine中使用的变量i将是循环变量i的内存地址，而不是i的值。
//因此，所有的goroutine都会读取相同的变量i。由于循环非常快，当goroutine开始运行时，循环已经结束，
//此时i的值被设置为5，因此所有goroutine都会输出Worker 5 starting。
//这样就无法实现我们想要的效果，即每个goroutine输出其自身的worker id。
//因此，必须在goroutine中创建一个新的变量来保存当前循环变量的值，以确保每个goroutine都可以独立地使用它们自己的worker id。

//这段代码创建了5个goroutine来模拟5个工人在同时开展工作。因为goroutine是并行执行的，所以它们会同时启动，并且输出顺序可能是随机的。

//在主函数中，我们定义了一个sync.WaitGroup类型的变量wg来等待所有goroutine执行完成。然后我们使用一个for循环来迭代5次，
//每次循环都会做以下事情：

//增加wg的计数器。
//在goroutine中使用新变量i保存当前循环变量i的值。
//启动一个匿名函数作为goroutine，在这个函数中调用worker函数并传递i作为参数。
//在goroutine结束时，减少wg的计数器。

//最后，我们调用wg.Wait()等待所有goroutine执行完成。当所有的goroutine都完成时，程序退出。每个工人的工作包括在启动时输出其worker id，
//然后休眠1秒钟来模拟它正在做某些工作，最后输出它已经完成了工作。
