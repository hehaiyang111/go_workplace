package main

import "fmt"

/**
接口2
*/

//type Describer interface {
//	describe()
//}
//
//type Person struct {
//	name string
//	age int
//}
//
//type Address_test15  struct {
//	name string
//	age int
//}

// 值接收器
//func (a Person)describe()  {
//	fmt.Println(a.name)
//}

// 指针接收器
//func (b *Address_test15)describe()  {
//	fmt.Println(b.name)
//}

//func main()  {
//	a := Person{
//		"hhy",
//		23,
//	}
//	b := Address_test15{
//		"Beijing Road",
//		115,
//	}
/**
第一种调用
*/
//a.describe()
//b.describe()

/**
第二种调用
1) 值接收器, 可以直接赋值给接口类型也可以把引用赋值给接口类型
2) 指针接收器, 只能将引用赋值给接口类型
*/
//	var add_copy Describer
//	add_copy = a
//	add_copy.describe()
//	add_copy = &b
//	add_copy.describe()
//}

/**
实现多个接口
*/
type SalaryCalculatorCopy interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

// 接口嵌套
type addTwoInterface interface {
	SalaryCalculatorCopy
	LeaveCalculator
}

type EmployeeCopy struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e EmployeeCopy) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e EmployeeCopy) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := EmployeeCopy{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	//e.DisplaySalary()

	var s SalaryCalculatorCopy = e
	s.DisplaySalary()

	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	//接口的嵌套
	var d addTwoInterface = e
	d.DisplaySalary()
}
