package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample") //os.CreateTemp(): 在指定目录中创建一个临时文件，返回该文件的文件对象和错误。
	check(err)

	fmt.Print("Temp file name: ", f.Name())

	defer os.Remove(f.Name()) //os.Remove(): 删除指定的文件或目录及其所有内容。如果操作成功，返回 nil；如果操作失败，返回错误 err。

	_, err = f.Write([]byte{1, 2, 3, 4}) //将字节切片写入文件
	check(err)

	dname, err := os.MkdirTemp("", "sampledir") //os.MkdirTemp(): 在指定目录中创建一个临时目录，返回该目录的名称和错误。
	check(err)
	fmt.Println("Temp dir name: ", dname)

	defer os.RemoveAll(dname) //os.RemoveAll(): 删除指定的文件或目录及其所有内容。如果操作成功，返回 nil；如果操作失败，返回错误 err。

	fname := filepath.Join(dname, "file1") //filepath.Join(): 将任意数量的路径元素放入单个路径中，会根据需要添加路径分隔符。
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
