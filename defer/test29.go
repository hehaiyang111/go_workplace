package main

import "fmt"

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
*/
type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		firstName: "hhy",
		lastName:  "zuishuai",
	}
	//用于延迟
	defer p.fullName()
	fmt.Printf("welcome ")
}
