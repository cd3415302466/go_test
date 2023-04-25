package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	//使用标准编码进行编码和解码
	sEnc := base64.StdEncoding.EncodeToString([]byte(data)) //base64.StdEncoding.EncodeToString(): 将字符串进行base64编码。
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc) //base64.StdEncoding.DecodeString(): 将字符串进行base64解码。
	fmt.Println(string(sDec))
	fmt.Println()

	//使用 URL 安全编码进行编码和解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data)) //base64.URLEncoding.EncodeToString(): 将字符串进行base64编码。
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc) //base64.URLEncoding.DecodeString(): 将字符串进行base64解码。
	fmt.Println(string(uDec))
}
