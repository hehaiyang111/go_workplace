package main

import "fmt"

/**
val,ok := <-a  使用变量ok可以检查信道是否关闭
*/
func producer(a chan int) {
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
}

func main() {
	a := make(chan int)
	go producer(a)
	for {
		val, ok := <-a
		if ok {
			fmt.Println("Revived ", val, ok)
		} else {
			break
		}
	}
}
