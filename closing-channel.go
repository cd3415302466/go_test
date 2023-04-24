package main

import "fmt"

func main() {
	// 创建一个可缓存大小为5的int类型的channel jobs 和一个bool类型的channel done
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 启动goroutine，从jobs channel中读取job值并处理
	go func() {
		for {
			j, more := <-jobs
			if more { // 如果channel未关闭，并且有新的job，则打印"received job"和job的值
				fmt.Println("received job", j)
			} else { // 如果channel已经关闭，则打印"received all jobs"，向done channel发送true值并退出goroutine
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 向jobs channel发送三个job，然后关闭channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 等待从done channel接收到true值，以确认所有job都已经被处理完毕
	<-done
}
