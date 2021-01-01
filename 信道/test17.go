package main

import "fmt"

func main() {
	//定义一些int类型的信道(chan) a
	var a chan int
	if a == nil {
		fmt.Println("channel is nil, please go to define it")
		// use make to define it
		a = make(chan int)
		fmt.Printf("Type of a is %T ", a)
	}
}
