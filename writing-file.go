package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n") //将字符串转换为字节切片
	err := os.WriteFile("./dat1", d1, 0644)
	check(err)

	f, err := os.Create("./dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} //将字符串转换为字节切片
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("%d bytes written\n", n2)

	n3, err := f.WriteString("writes\n") //将字符串转换为字节切片
	check(err)
	fmt.Printf("%d bytes written\n", n3)

	f.Sync() //将缓冲区的数据写入磁盘

	w := bufio.NewWriter(f) //创建一个新的Writer对象
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("%d bytes written\n", n4)

	w.Flush() //将缓冲区的数据写入磁盘
}
