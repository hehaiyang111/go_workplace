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

func server1(a chan string) {
	time.Sleep(6 * time.Second)
	a <- "from server1"
}

func server2(b chan string) {
	time.Sleep(3 * time.Second)
	b <- "from server2"
}

func main() {
	a := make(chan string)
	b := make(chan string)
	go server1(a)
	go server2(b)
	select {
	case s1 := <-a:
		fmt.Println(s1)
	case s2 := <-b:
		fmt.Println(s2)
	}
}
