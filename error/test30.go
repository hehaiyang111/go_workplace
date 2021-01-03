package main

import (
	"fmt"
	"net"
)

/**
错误表示了程序中出现了异常情况，比如当我们试图打开一个文件时，文件系统却并没有这个文件，这就是异常情况，它用一个错误来表示
在Go中，错误一直是很常见的。错误用内建的error类型表示。
*/
/**
Example1：程序试图再开一个并不存在的文件
按照Go的管理，在处理错误的时候，通常都是将返回的错误与nil比较。nil值表示了没有错误发生，而非nil值表示出现了错误。
*/
//func main() {
//	f, err := os.Open("./test.txt")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(f.Name(), "opened successfully")
//}

/**
Example2：从错误获取更多信息的不同方法
method: 对底层类型进行断言，然后通过调用该结构体类型的方法，来获取更多的信息
*/
// 标准库中的DNSError结构体类型定义如下：
//type DNSError struct {
//	...
//}
//
//func (e *DNSError) Error() string {
//	...
//}
//func (e *DNSError) Timeout() bool {
//	...
//}
//func (e *DNSError) Temporary() bool {
//	...
//}
func main() {
	addr, err := net.LookupHost("www.liuxinyi.top")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}
