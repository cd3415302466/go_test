package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"} // 切片
	sort.Strings(strs)              // 字符串排序
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4} // 切片
	sort.Ints(ints)        // 整型排序
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints) // 判断是否已经排序
	fmt.Println("Sorted: ", s)
}
