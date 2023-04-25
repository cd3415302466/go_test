package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) //strconv.ParseFloat(): 将字符串转换为浮点型。
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) //strconv.ParseInt(): 将字符串转换为有符号整型。
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) //strconv.ParseUint(): 将字符串转换为无符号整型。
	fmt.Println(u)

	k, _ := strconv.Atoi("135") //strconv.Atoi(): 将字符串转换为int类型。
	fmt.Println(k)

	_, e := strconv.Atoi("wat") //strconv.Atoi(): 将字符串转换为int类型。
	fmt.Println(e)
}
