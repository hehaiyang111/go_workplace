package main

import (
	"fmt"
	"time"
)

/**
panic and recover
Explain:对于大多种情况error就够用了,但是有些情况下程序发生异常时无法继续执行.在这种情况下我们会使用panic终止程序.当函数发生panic时候他会终止运行
相当于try-catch-finally
一般不用。
注意：应该尽可能的使用错误，而不是使用panic和recover。只有当程序不能继续执行的时候，才应该使用panic和recover机制
*/
/**
panic有两个合理用例：
1. 发生了一个不能恢复的错误：此时程序不能继续运行。例如web服务器无法绑定所要求的端口。在这种情况下，就应该使用panic，因为如果不能绑定端口，啥也做不了。
2. 发生了一个编程上的错误：加入我们有一个接受指针参数的方法，而其他人使用nil作为参数调用了它。在这种情况下，我们可以使用panic，因为这是一个编程错误：用nil参数调用了一个只能接受合法指针的方法。
*/
/**
Example1
*/
/*
func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func main() {
	firstName := "hehaiyang"
	fullName(&firstName,nil)
	fmt.Println("returned normally from main")
}
*/

/**
Example2：发生panic时的defer
Explain：如果有延迟函数，会先调用它，然后程序控制返回到函数调用方
比上面的程序多了两个defer
*/
/*
func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
*/

/*
	recover是一个内建函数，用于重新获得panic协程的控制
	用法：只有在延迟函数的内部，调用recover才有用。在延迟函数内调用recover，可以提取到panic的错误信息。并且停止panic续发时间，程序运行恢复正常。如果在延迟函数的外部调用recover，就不能停止panic续发事件。
*/
/*
func fullName(firstName *string, lastName *string)  {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName() {
	if r := recover(); r !=nil {
		fmt.Println("recovered from ", r)
	}
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
*/
/**
只有在相同的Go协程中调用recover才管用。recover不能恢复一个不同协程的panic.我们来举例一下。
*/
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}
func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}
func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}
func main() {
	a()
	fmt.Println("normally returned from main")
}
