package main

import "fmt"

/**
接口：
1. 接口定义一个对象的行为
2. 接口制定了一个借口需要实现的方法，但是具体如何实现不管。
*/

/**
计算不同类型工人的工资
*/

type SalaryCalculator interface {
	CalculateSalary() int
}

//长期工
type Permanent struct {
	empId, basicPay, pf int
}

//短期工
type Contract struct {
	empId, basicPay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func calculatorSum(v []SalaryCalculator) {
	sum := 0
	for _, value := range v {
		sum = sum + value.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", sum)
}
func main() {
	p := Permanent{
		1,
		500,
		2000,
	}
	c := Contract{
		2,
		500,
	}
	// p和c都是接口SalaryCalculator类型的。 所以创建一个这个类型的数组
	employee := []SalaryCalculator{p, c}

	calculatorSum(employee)

}

/**
可接受任何参数
*/
//func describe(i interface{}) {
//	fmt.Printf("Type = %T, value = %v\n", i, i)
//}

/**
类型断言
*/
//func assert(i interface{}) {
//	v, ok := i.(int)
//	fmt.Println(v, ok)
//}
//func main() {
//	var s interface{} = 56
//	assert(s)
//	var i interface{} = "Steven Paul"
//	assert(i)
//}
// Output： 56 true  \n     0 false

/**
switch 类型
*/

//func findType(i interface{}) {
//	switch i.(type) {
//	case string:
//		fmt.Printf("I am a string and my value is %s\n", i.(string))
//	case int:
//		fmt.Printf("I am an int and my value is %d\n", i.(int))
//	default:
//		fmt.Printf("Unknown type\n")
//	}
//}
//func main() {
//	findType("Naveen")
//	findType(77)
//	findType(89.98)
//}

/**
如果一个类型实现了接口， 那么该类型与其实现的接口就可以相互比较
*/
//type Describer interface {
//	Describe()
//}
//type Person struct {
//	name string
//	age  int
//}
//
//func (p Person) Describe() {
//	fmt.Printf("%s is %d years old", p.name, p.age)
//}
//
//func findType(i interface{}) {
//	switch v := i.(type) {
//	case Describer:
//		v.Describe()
//	default:
//		fmt.Printf("unknown type\n")
//	}
//}
//
//func main() {
//	findType("Naveen")
//	p := Person{
//		name: "Naveen R",
//		age:  25,
//	}
//	findType(p)
//}
