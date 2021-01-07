package main

import "fmt"

/*
	匿名函数
*/

/*
	Example：首先编写一个简单的例子，把函数赋值给一个变量
*/
/*
func main() {
	a := func() {
		// 把没有名字的函数赋值给变量a
		fmt.Println("I am a func hhhh ~")
	}
	fmt.Printf("%T", a)
}
*/

/*
	Example1: 要调用一个匿名函数，可以不用赋值给变量
*/
/*
func main() {
	func() {
		fmt.Println("hello world first class function")
	}()
}
*/

/**
Example1: 要调用一个匿名函数，可以不用赋值给变量,还可以传参
*/
/*
func main() {
	func(a string) {
		fmt.Println("Welcome", a)
	}("hhy")
}
*/

/*
	Example2:自定义函数类型
*/
/*
type add func(a int, b int) int

func main() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5,6)
	fmt.Println("Sum", s)
}
*/

/*
	高阶函数，满足下列条件之一
	1. 接受一个或多个函数作为参数
	2. 返回值是一个函数
*/
/*
func simple(a func(a,b int) int) {
	fmt.Println(a(60,7))
}
func main() {
	f := func(a,b int) int{
		return a + b
	}
	simple(f)
}
*/

/*
	Example3:在其他函数中返回函数
*/
/*
func simple() func(a, b int) int {
	f := func(a, b int) int{
		return a + b
	}
	return f
}
func main() {
	s := simple()
	fmt.Println(s(60, 7))
}
*/

/**
闭包:变量存在函数体的外部
result:
Hello World
Hello Everyone
Hello World Gopher
Hello Everyone !
*/
func appendStr() func(a string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
