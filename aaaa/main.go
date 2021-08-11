package main

import (
	"fmt"
	"path"
)

func main() {
	files := "test.txt"
	fmt.Println(path.Ext(files)) //获取路径中的文件的后缀 .txt

}
