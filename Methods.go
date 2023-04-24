package main

import "fmt"

// 方法 method函数

type rect struct {
	width, height int
}

//定义了一个名为 rect 的结构体类型，它具有两个整数字段：width 和 height。

func (r *rect) area() int {
	return r.width * r.height
}

//area() 方法计算并返回矩形的面积，即将 width 和 height 相乘。

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

//perim() 方法计算并返回矩形的周长，即将 width 和 height 分别乘以 2 后相加。

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	//创建了指向 r 的指针 rp，然后调用 area() 和 perim() 方法。
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	//通过值和指针两种方式调用 rect 对象的方法，结果都是相同的
}
