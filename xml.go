package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"` //attr 表示这是一个属性而不是一个元素。
	Name    string   `xml:"name"`    //name 元素拥有一个内部的, chardata 注释，因此解析器将在遇到字符数据时将其作为 Name 字段的值。
	Origin  []string `xml:"origin"`  //origin 元素拥有一个 innerxml 注释，因此解析器将在遇到该元素时将其未解析的原始字节片段作为 Origin 字段的值。

}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ") //MarshalIndent 函数将 Go 值编码为 XML 字符串。
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil { //Unmarshal 函数将 XML 字符串解码为 Go 值。
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`            //XMLName 字段的类型必须为 xml.Name，且它必须有一个非空的本地名称或空格分隔的名称列表。
		Plants  []*Plant `xml:"parent>child>plant"` //parent>child>plant 元素路径告诉编码器将每个 Plant 实例编码为 <parent><child><plant> 元素。

	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
