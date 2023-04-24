package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

//geometry 接口

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	//main 函数中，首先创建了一个矩形对象 r 和一个圆形对象 c。然后调用了 measure 函数来测量这两个对象的面积和周长。
	//measure 函数接受一个 geometry 类型的参数，并调用其相应的方法来计算面积和周长，并将结果输出到控制台。
}

//通过这个示例程序，我们可以看到 Go 语言中的接口和结构体的基本用法，以及如何使用它们来实现多态性和抽象。
