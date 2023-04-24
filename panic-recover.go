package main

import "os"

func main() {
	defer func() {
		if r := recover(); r != nil {
			// 处理 panic 异常
			println("Recovering from panic:", r)
		}
	}()

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	println("This line will never be executed.")
}
