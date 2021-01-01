package main

import (
	"fmt"
	"time"
)

/**
select语句用于在发送/接受信道操作中进行选择。
select语句会一直堵塞，直到发送/接受操作准备就绪。如果有多个信道操作准备完毕，select会随机的选取其中之一执行。
该语法与switch相似，不同的是，这里的每个case语句都是信道操作。
show my code
*/
//
//func server1(a chan string) {
//	time.Sleep(6 * time.Second)
//	a <- "from server1"
//}
//
//func server2(b chan string) {
//	time.Sleep(3 * time.Second)
//	b <- "from server2"
//}
//
//func main() {
//	a := make(chan string)
//	b := make(chan string)
//	go server1(a)
//	go server2(b)
//	select {
//	case s1 := <-a:
//		fmt.Println(s1)
//	case s2 := <-b:
//		fmt.Println(s2)
//	}
//}

/**
select应用：
目的:假设我们有一个关键性应用，需要尽快的输出返回给用户
需求：有一个应用的数据库赋值并且存储在世界各地的服务器上。假设函数server1和server2与这样不同区域的两台服务器进行通信。每台服务器的负载和网络延时决定了他的响应时间
分析：我们像两台服务器发送请求，并且使用select语句等待相信的信道发出相应。select会选择优先相应的服务器，而忽略其他的相应。使用这种方法，我们可以向多个服务器发送请求，并给返回给用户最快的响应的那台服务器
*/
//func process(a chan string) {
//	time.Sleep(10 * time.Second)
//	a <- "tomorrow will be ok"
//}
//func main() {
//	a := make(chan string)
//	go process(a)
//	for {
//		// 模仿一下耗时，要不然for循环执行很快的
//		time.Sleep(1 * time.Second)
//		select {
//		case v := <-a:
//			fmt.Println("received value: ", v)
//			//如果有输入就return退出for循环
//			return
//			//默认情况下
//		default:
//			fmt.Println("no value received")
//		}
//	}
//}

/**
死锁与默认情况
死锁:当我们创建一个信道存放在select,试着去读取他的值,但是没有go协程像该信道输入值写入数据,select语句就是一直阻塞,导致死锁.
默认情况:如果存在默认情况,就不会发生死锁.因为在没有其他case准备就绪时,会执行默认情况
*/
//func main()  {
//	a := make(chan string)
//	select {
//	case <-a:
//		fmt.Println("no deadlock")
//	default:
//		fmt.Println("default case executed")
//	}
//}

/**
随机选择:当两个go协程都满足条件时,select会随机选取一个case执行
*/
func server1(a chan string) {
	a <- "go server1"
}

func server2(b chan string) {
	b <- "go server2"
}

func main() {
	a := make(chan string)
	b := make(chan string)
	go server1(a)
	go server2(b)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-a:
		fmt.Println(s1)
	case s2 := <-b:
		fmt.Println(s2)
	}
}
