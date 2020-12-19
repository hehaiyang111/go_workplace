package main

import "fmt"

/**
	结构体
 */

type Employee struct {
	firstName, SecondName string
	age, salary int
}

type Address struct {
	addressName string
	addressAge int
}

type Student struct {
	name string
	age int
	address Address
}

func main()  {
	/**
		用法1
	 */
	//a := Employee{
	//	"Hehehe","Shiniba",18,20}
	//
	//b := Employee{
	//	"Hehehe","Shinima",30,20}
	//fmt.Println(a.age)
	//fmt.Println(b.age)

	/**
		用法2
	 */
	//var a Employee
	//var b Employee
	//a.age = 3
	//b.age = 4
	//fmt.Println(a.age)
	//fmt.Println(b.age)

	/**
		带有引用参数
	 */
	a := Student{name: "hhy", age: 23, address: Address{addressName: "BeiJingRoad",addressAge: 118}}

	fmt.Println(a.address.addressName)
	fmt.Print(a.address.addressAge)



}