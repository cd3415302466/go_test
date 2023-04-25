package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// check 函数用于检查错误并在发生错误时 panic。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 创建名为 subdir 的目录，权限为 0755。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 在 main 函数结束时，删除名为 subdir 的目录及其所有内容。
	defer os.RemoveAll("subdir")

	// createEmptyFile 函数用于创建指定名称的空文件。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	// 在名为 subdir 的目录下创建名为 file1 的空文件。
	createEmptyFile("subdir/file1")

	// 在名为 subdir/parent/child 的目录下递归创建目录。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// 在名为 subdir/parent 目录下创建名为 file2、file3 和 subdir/parent/child/file4 的空文件。
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// 读取名为 subdir/parent 的目录中的条目，并打印每个条目的名称和是否为目录。
	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 更改当前工作目录为 subdir/parent/child，并读取该目录中的条目。
	err = os.Chdir("subdir/parent/child")
	check(err)
	c, err = os.ReadDir(".")
	check(err)

	// 打印 subdir/parent/child 目录中的每个条目的名称和是否为目录。
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 更改当前工作目录为 subdir，然后递归遍历名为 subdir 的目录及其所有子目录，并使用 visit 函数处理每个文件或目录。
	err = os.Chdir("../../..")
	check(err)
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// visit 函数用于处理 filepath.Walk 遍历的每个文件或目录。
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// 打印文件或目录的路径和是否为目录。
	fmt.Println(" ", p, info.IsDir())
	return nil
}
