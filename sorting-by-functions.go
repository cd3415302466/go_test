package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s) // 返回切片的长度
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // 交换切片的元素
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j]) // 比较切片的元素的长度
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"} // 创建一个字符串切片 fruits 包含 "peach", "banana", "kiwi"
	sort.Sort(byLength(fruits))                   // 将 fruits 转换为 byLength 类型，再通过 sort.Sort 方法排序
	fmt.Println(fruits)                           // 将 fruits 转换为 byLength 类型，再通过 sort.Sort 方法排序
}
