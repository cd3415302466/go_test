package main

import "fmt"

type base struct {
	num int
}

// 定义结构体类型 base

func (b base) describe() string { //base 类型有一个 num 字段和一个 describe() 方法，该方法返回一个字符串，描述该结构体的属性值。
	return fmt.Sprintf("base with num %v", b.num)
}

//定义接口类型 describer

type container struct { //container 类型有一个 base 字段和一个 str 字段。base 字段的类型是 base，str 字段的类型是 string。
	base
	str string
}

// 定义结构体类型container

func main() {
	co := container{
		base: base{ //base 类型有一个 num 字段和一个 describe() 方法，该方法返回一个字符串，描述该结构体的属性值。
			num: 1, //num 字段的值为 1。
		},
		str: "some name", //str 字段的值为 "some name"。
	}
	//container 类型有一个 base 字段和一个 str 字段。base 字段的类型是 base，str 字段的类型是 string。

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co                   //container 类型的值可以赋值给 describer 类型的变量。
	fmt.Println("describe:", d.describe()) //d.describe() 调用的是 container 类型的方法，而不是 base 类型的方法。
}
