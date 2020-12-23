package main

import (
	"fmt"
	"time"
)

/**
信道没有赋值的时候是处于阻塞状态
time.Sleep(2 * time*Second): 等待两秒以后再给信道赋值，可以看到没有赋值的时候，主进程堵塞。
*/
func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(2 * time.Second)
	done <- true
}
func main() {
	done := make(chan bool)
	// 执行协程
	go hello(done)
	<-done
	fmt.Printf("main function")
}
