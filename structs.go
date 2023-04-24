package main

import "fmt"

//结构

type person struct {
	name string
	age  int
}

//定义了一个名为 person 的结构体，包含两个字段 name 和 age

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
	//使用&操作符获取p的地址，并将其作为指向person类型的指针返回。
}

// 定义了一个 newPerson 函数，用于创建并返回一个指向 person 结构体的指针。

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))
	//将创建一个指向person类型的指针，并且这个person对象的name属性为"Jon"

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
