package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5)

	//创建一个整数切片并将其作为参数调用sum函数，通过在切片名称后面加上省略号（...）来将切片解包为单独的参数。
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum(nums...)
}
