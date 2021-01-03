package main

import (
	"fmt"
	"sync"
)

/**
defer：含有defer语句的函数，会在函数将要返回之前，调用另外一个函数。
*/

//func finished() {
//	fmt.Println("program finished")
//}
//func largest(nums []int) {
//	// defer后的方法，在当前方法largest执行完成后执行
//	defer finished()
//	fmt.Println("program started")
//	max := nums[0]
//	for _,val  := range nums {
//		if val > max {
//			max = val
//		}
//	}
//	fmt.Println("Largest number in", nums, "is", max)
//}
//
//func main() {
//	a := []int{123,432,2134,3421,5421,234}
//	largest(a)
//}

/**
延迟方法：不仅可以用于函数的调用，同样也可以用于方法的延迟
//*/
//type person struct {
//	firstName string
//	lastName  string
//}
//
//func (p person) fullName() {
//	fmt.Printf("%s %s", p.firstName, p.lastName)
//}
//
//func main() {
//	p := person{
//		firstName: "hhy",
//		lastName:  "zuishuai",
//	}
//	//用于延迟
//	defer p.fullName()
//	fmt.Printf("welcome ")
//}

/**
实参取值，在Go语言中，并非在调用延迟函数的时候才确定实参，而是当执行defer语句的时候，就会对延迟函数的实参进行求值。
*/
//func printA(a int) {
//	fmt.Println("value of a in deferred function", a)
//}
//func main() {
//	a := 5
//	defer printA(a)
//	a = 10
//	fmt.Println("value of a before deferred function call", a)
//}
// result: value of a before deferred function call 10
//         value of a in deferred function 5

/**
defer栈，当一个函数内多次调用defer时，Go会把defer调用放入一个栈中	，随后按照后进先出的顺序执行。
*/
//func main() {
//	name := "hehaiyang"
//	fmt.Printf("Orignal String: %s\n", string(name))
//	for _, v := range []rune(name) {
//		defer fmt.Printf("%c", v)
//	}
//}

/**
defer用途
通过未使用defer和使用defer的两段代码做对比。
*/
// 未使用
//type rect struct {
//	length int
//	width  int
//}
//
//func (r rect) area(wg *sync.WaitGroup) {
//	if r.length < 0 {
//		fmt.Printf("rect %v's length should be greater than zero\n", r)
//		wg.Done()
//		return
//	}
//	if r.width < 0 {
//		fmt.Printf("rect %v's width should be greater than zero\n", r)
//		wg.Done()
//		return
//	}
//	area := r.length * r.width
//	fmt.Printf("rect %v's area %d\n", r, area)
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//	r1 := rect{-67, 89}
//	r2 := rect{5, -67}
//	r3 := rect{8, 9}
//	rects := []rect{r1, r2, r3}
//	for _, v := range rects {
//		wg.Add(1)
//		go v.area(&wg)
//	}
//	wg.Wait()
//	fmt.Println("All go routines finished executing")
//}

//使用defer
type rect struct {
	width  int
	length int
}

func (r rect) area(wg *sync.WaitGroup) {
	// 只有在area函数返回的时候才会调用
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
