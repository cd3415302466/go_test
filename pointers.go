package main

import "fmt"

// 该程序定义了两个函数 zeroval 和 zeroptr，它们分别接收一个整型参数 ival 和一个整型指针参数 iptr。
// 其中 zeroval 函数通过值传递方式将参数 ival 的值设置为 0，而 zeroptr 函数通过指针传递方式将参数指针 iptr 所指向的变量的值设置为 0。
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

//在 main 函数中，定义了一个整型变量 i 并初始化为 1，然后依次调用了 zeroval、zeroptr 函数，并输出了 i 的值以及其地址。
//其中，在调用 zeroval 函数后，i 的值仍为 1，因为 zeroval 函数只是修改了 ival 这个局部变量的值，而没有修改实参 i 的值。
//而在调用 zeroptr 函数后，i 的值被成功修改为 0，因为 zeroptr 函数通过指针传递方式修改了实参 i 的值。
//最后，程序输出了 i 的地址，即 &i，这个值是一个指向 i 变量的指针。

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
