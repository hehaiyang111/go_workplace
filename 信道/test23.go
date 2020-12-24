package main

import "fmt"

/**
重写一下test19
1. func digit：把每一位的值取出来放到一个信道里
2. func calculateCubes_test：遍历信道里的值，进行计算求立方和，然后将值赋值给另一个信道
3. func calculateSquares_test：遍历信道里的值，进行计算求平方和，然后将值赋值给另一个信道
*/

func digit(a int, b chan int) {
	for a != 0 {
		c := a % 10
		b <- c
		a /= 10
	}
	close(b)
}

func calculateCubes_test(a int, b chan int) {
	sum := 0
	c := make(chan int)
	go digit(a, c)
	for val := range c {
		sum += val * val * val
	}
	b <- sum
}

func calculateSquares_test(a int, b chan int) {
	sum := 0
	c := make(chan int)
	go digit(a, c)
	for val := range c {
		sum += val * val
	}
	b <- sum
}

func main() {
	a := 589
	b := make(chan int)
	c := make(chan int)
	go calculateSquares_test(a, b)
	go calculateCubes_test(a, c)
	d, e := <-b, <-c
	fmt.Println("Final output", d+e)
}
