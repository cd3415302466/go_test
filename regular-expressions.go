package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //MatchString 函数检查字符串是否匹配表达式，这里检查字符串是否以 p 开头，以 ch 结尾，中间是任意的一个或多个字符。
	//如果匹配成功，返回 true，否则返回 false。
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") //Compile 函数将字符串形式的正则表达式编译成正则表达式对象。

	fmt.Println(r.MatchString("peach")) //MatchString 函数检查字符串是否匹配表达式，这里检查字符串是否以 p 开头，以 ch 结尾，中间是任意的一个或多个字符。

	fmt.Println(r.FindString("peach punch")) //FindString 函数返回字符串中第一次匹配成功的字符串。

	fmt.Println("idx:", r.FindStringIndex("peach punch")) //FindStringIndex 函数返回字符串中第一次匹配成功的字符串的开始和结束位置。

	fmt.Println(r.FindStringSubmatch("peach punch")) //FindStringSubmatch 函数返回字符串中第一次匹配成功的字符串和匹配成功的字符串的子匹配。

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //FindStringSubmatchIndex 函数返回字符串中第一次匹配成功的字符串和匹配成功的字符串的子匹配的开始和结束位置。

	fmt.Println(r.FindAllString("peach punch pinch", -1)) //FindAllString 函数返回字符串中所有匹配成功的字符串。

	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //FindAllStringSubmatchIndex 函数返回字符串中所有匹配成功的字符串和匹配成功的字符串的子匹配的开始和结束位置。

	fmt.Println(r.FindAllString("peach punch pinch", 2)) //FindAllString 函数返回字符串中所有匹配成功的字符串。

	fmt.Println(r.Match([]byte("peach"))) //Match 函数检查字节数组是否匹配表达式。

	r = regexp.MustCompile("p([a-z]+)ch") //MustCompile 函数将字符串形式的正则表达式编译成正则表达式对象，如果成功返回正则表达式对象，否则 panic。
	fmt.Println("regexp:", r)             //regexp: p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //ReplaceAllString 函数将匹配到的字符串替换成指定的字符串。

	in := []byte("a peach")                    //ReplaceAllFunc 函数将匹配到的字符串替换成指定的字符串。
	out := r.ReplaceAllFunc(in, bytes.ToUpper) //ReplaceAllFunc 函数将匹配到的字符串替换成指定的字符串。
	fmt.Println("out:", string(out))           //out: a PEACH
}
