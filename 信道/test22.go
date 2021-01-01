package main

import "fmt"

/**
for range可以用于在信道关闭之前，从信道接受数据
*/

func producer_test(a chan int) {
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
}

func main() {
	a := make(chan int)
	go producer_test(a)
	for val := range a {
		fmt.Println("Received ", val)
	}
}
