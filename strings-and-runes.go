package main

import (
	"fmt"
	"unicode/utf8" //Go 语言的标准库 unicode/utf8 来处理 Unicode 字符串
)

func main() {
	const s = "hello world" //定义了一个常量 s，其值为 "hello world"。

	fmt.Println("len:", len(s))

	//使用 len() 函数来获取字符串 s 的字节数，并将其打印到控制台上。由于使用的是 UTF-8 编码，因此每个非 ASCII 字符会占用多个字节，
	//所以在这里使用 len() 函数无法直接得到字符串中字符的数量。

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i]) //for 循环遍历字符串 s 中的每个字节，并分别打印它们的十六进制表示方式（使用 %x 格式化字符串）。
	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s))
	//使用 utf8.RuneCountInString() 函数来获取字符串 s 中的字符数，
	// 并将其打印到控制台上。这个函数返回的是字符串中 Unicode 字符的数量，而不是字节数。

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}
	//使用 range 关键字和 for 循环遍历字符串 s 中的每个 Unicode 字符。在循环体中，
	//使用 fmt.Printf() 函数和 %#U 格式化字符串来打印当前字符的 Unicode 码点及其在字符串中的索引位置。
	fmt.Println("\nUsing range to print runes:")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		//程序使用 utf8.DecodeRuneInString() 函数来遍历字符串并解码每个 Unicode 字符。在循环体中，首先使用当前索引位置调用该函数，
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
		//并将返回值分别存储在 runeValue 和 width 变量中。然后打印当前字符的 Unicode 码点及其在字符串中的索引位置，
		examineRune(runeValue)
		//并调用 examineRune() 函数检查该字符是否为 't' 或者 'l'。循环变量 i 每次增加 width，即当前字符所占用的字节数，
		//以便继续遍历下一个字符。
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'l' {
		fmt.Println("found so sua")
	}
}

//调用了 examineRune() 函数，第一个参数是当前遍历到的 Unicode 字符。在这个函数中，
//使用简单的条件语句判断字符是否为 't' 或者 'l'，如果是，则在控制台上打印相应的消息。
