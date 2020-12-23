package main

import "fmt"

/**
避免信道死锁  当go协程给信道发送数据时 按道理得有go协程来接受数据
*/

// 死锁
//func main()  {
//	ch := make(chan int)
//	ch <- 5
//}

/**
单向信道
*/
// 把双向信道可转为唯送和唯收信道
func sendDate(a chan<- int) {
	a <- 10
}

func main() {
	// 建立一个双向信道
	a := make(chan int)
	// 信道要和协程一起用哦~
	go sendDate(a)
	fmt.Print(<-a)
}
