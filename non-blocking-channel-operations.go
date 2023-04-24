package main

import "fmt"

func main() {
	// 创建两个 channel: messages 和 signals
	messages := make(chan string)
	signals := make(chan bool)

	// 在第一个 select 语句中尝试从 messages channel 接收消息，
	// 如果没有则启动一个 goroutine 发送一条消息到该 channel
	go func() {
		messages <- "hello"
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 向 messages channel 发送一条消息，并在第二个 select 语句中接收该消息
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在第三个 select 语句中尝试从 messages 或 signals channel 中接收消息，
	// 如果都没有消息则执行 default 分支中的代码块
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
