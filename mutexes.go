package main

import (
	"fmt"
	"sync"
)

// Container 表示带有关联互斥锁的线程安全映射容器。
type Container struct {
	mu        sync.Mutex
	container map[string]int
}

// inc 方法以线程安全的方式将给定键相关的值加1。
func (c *Container) inc(name string) {
	// 加锁确保同一时间只有一个 goroutine 访问映射容器。
	c.mu.Lock()
	defer c.mu.Unlock()

	// 将给定键相关的值加1。
	c.container[name]++
}

func main() {
	// 创建新的 Container 实例，初始化为包含两个键 "a" 和 "b" 的映射容器，初始值均为0。
	c := Container{
		container: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// 定义一个函数用于将与给定键相关的值增加 n 次。
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 启动三个 goroutine，分别将与键 "a" 和 "b" 相关的值增加10000次。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("a", 10000)

	// 等待所有三个 goroutine 执行完成。
	wg.Wait()

	// 打印修改后的映射容器。
	fmt.Println(c.container)
}
