package main

import "fmt"

/**
	引用
 */

func a_b(val *int) {
	c := 7
	val = &c
	fmt.Print(val)
}

func main() {
	b := 2
	a := &b
	fmt.Println(a)
	a_b(a)
	//fmt.Print(*a)
}
