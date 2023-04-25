package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 niubi" //字符串
	h := sha256.New()   //New 函数返回一个新的使用SHA256校验和的hash.Hash接口。
	h.Write([]byte(s))  //Write 函数将数据写入哈希接口。

	bs := h.Sum(nil) //Sum 函数返回一个表示当前哈希的切片。

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
