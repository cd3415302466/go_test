package main

import "os"

func main() {
	panic("a problem")
	//这段代码会导致一个 panic 异常，因为在 panic("a problem") 之后的代码都不会被执行。
	//在发生 panic 时，程序会立即终止，并且将错误信息输出到控制台。

	_, err := os.Create("/tmp/file")
	//在这个例子中，即使创建 /tmp/file 文件的代码正常执行，也永远不会被执行。
	//如果您想要在程序中处理错误并继续执行代码，则应该使用 defer 和 recover 来恢复 panic。
	if err != nil {
		panic(err)
	}
}
