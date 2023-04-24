package main

import "fmt"

func MapKeys[k comparable, v any](m map[k]v) []k { // MapKeys 函数的参数 m 的类型是 map[k]v，返回值的类型是 []k。
	r := make([]k, 0, len(m)) //r 的类型是 []k。
	for k := range m {        //遍历 m 的键。
		r = append(r, k) //将键添加到 r 中。
	}
	return r //返回 r。
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) { // Push 方法的参数 v 的类型是 T。
	if lst.tail == nil { //如果 lst.tail 为 nil，说明链表为空。
		lst.head = &element[T]{val: v} //创建一个新的元素，将其赋值给 lst.head。
		lst.tail = lst.head            //将 lst.tail 指向 lst.head。
	} else {
		lst.tail.next = &element[T]{val: v} //创建一个新的元素，将其赋值给 lst.tail.next。
		lst.tail = lst.tail.next            //将 lst.tail 指向 lst.tail.next。
	}
}

func (lst *List[T]) GetAll() []T { // GetAll 方法的返回值的类型是 []T。
	var elems []T                             //elems 的类型是 []T。
	for e := lst.head; e != nil; e = e.next { //遍历链表。
		elems = append( //将元素的值添加到
			elems, //elems 中。
			e.val, //e.val 的类型是 T。
		)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.GetAll())

}
