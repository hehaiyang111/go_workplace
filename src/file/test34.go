package main

import (
	"fmt"
	"os"
)

/*
	将整个文件读取到内存
		* 使用绝对路径
		* 使用命令行标记来传递文件路径
		* 将文件绑定在二进制文件中
*/
/*
func main() {
	// 使用绝对路径
	data, err := ioutil.ReadFile("D:\\go_workplace\\src\\file\\test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
*/

/*
	按字节读取
*/
func main() {
	filePath := "D:\\go_workplace\\src\\file\\test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("has some error")
	}
	//a := make([]byte,0)
	b := make([]byte, 3)
	// TODO 用缓冲区读取
	for {
		f.Read(b)
	}
}
