package main

import "fmt"

/**
之前讨论的是无缓冲信道，无缓冲信道的发送和接受是阻塞的
缓冲信道：只在缓冲满的时候，才会阻塞向缓冲信道发送数据
*/

/**
Example1：缓冲区的定义
*/
//func main() {
//	//使用make定义一个capacity为2的信道
//	a := make(chan string, 2)
//	a <- "s"
//	a <- "d"
//	fmt.Println(<-a)  		//读取完以后缓冲区就空了
//	fmt.Println(<-a)
//
//
//}

/**
Example：理解缓冲区
*/
//func readNumber(a chan int){
//	for i:=0; i<5; i++ {
//		a <- i
//		fmt.Println("successfully wrote", i, "to number")
//	}
//	close(a)
//}
//func main() {
//	a := make(chan int, 2)
//	go readNumber(a)
//	time.Sleep(2 * time.Second)
//	for b:= range a {
//		fmt.Println("read value", b,"from a")
//		// 读一个释放1个， go下称就会赋值一下，执行一下33行
//		time.Sleep(2 * time.Second)
//	}
//}

/**
Example：死锁
*/
func main() {
	a := make(chan int, 2)
	a <- 2
	a <- 3
	a <- 4 // 这个时候发生阻塞，需要其他协程进行释放。可惜没有，死锁了
	fmt.Println(<-a)
	fmt.Println(<-a)
}
