package main

import "fmt"

/**
利用到了 协程和 信道
计算一个数每一位的平方和与立方和，然后把平方和与立方和相加
*/

/**
1. 计算平方和的协程
	params: int
			chan int
*/
func calculateSquares(a int, b chan int) {
	sum := 0
	for a != 0 {
		c := a % 10
		sum += c * c
		a /= 10
	}
	b <- sum
}

/**
2. 计算立方和的协程
*/
func calculateCubes(a int, d chan int) {
	sum := 0
	for a != 0 {
		c := a % 10
		sum += c * c * c
		a /= 10
	}
	d <- sum
}

func main() {
	a := 589
	b := make(chan int)
	d := make(chan int)

	go calculateCubes(a, b)
	go calculateSquares(a, d)
	g, f := <-b, <-d

	fmt.Print("final output  ", g+f)
}
