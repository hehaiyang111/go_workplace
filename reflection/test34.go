package main

import (
	"fmt"
	"reflect"
)

/*
	什么是反射？
	反射就是程序能够在运行时检查变量和值，求出他们的类型。
*/

/*
	problem1:为什么要检查变量，确定变量的类型呢？？
/*

/*
	如果程序中每个变量都是我们自己定义的，那么在编译时就可以知道变量类型了，大多数程序是这样，但并非总是如此。
*/
/*
func main() {
	i := 10
	fmt.Printf("%d %T", i, i)
}
*/

/**
把结构体作为参数，并用它创建一个SQL
*/
/*
type order struct {
	ordId      int
	customerId int
}

func createSQL(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func main() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createSQL(o))
}
*/

/**
在Go语言中，reflect实现了运行时反射。reflect包会帮助识别interface{}变量的底层具体类型和具体值。createQuery函数接受interface{}参数，根据它的具体类型和具体值，创建sql查询
先了解一下reflect包中的几种类型方法
reflect.Type表示interface{}的具体类型，而reflect.Value表示他的具体值
reflect.TypeOf()和reflect.ValueOf()两个函数可以分别返回reflect.Type和reflect.Value
*/

/*
type order struct {
	ordId      int
	customerId int
}
func createSQL(q interface{}) {
	t := reflect.TypeOf(q)
	//t := reflect.TypeOf(q).Name()
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createSQL(o)
}
*/

/*
	reflect包中还有一个重要的类型：Kind
	在反射包中，Kind和Type的类型可能看起来很相似，但是在下面可以看出他们的区别
*/
/*
type order struct {
	ordId      int
	customerId int
}
func createSQL(q interface{}) {
	t := reflect.TypeOf(q)
	v := t.Kind()
	fmt.Println("Type ", t) //order
	fmt.Println("Kind ", v)	//order具体地类型struct
}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createSQL(o)
}
*/

/*
	NumField():方法返回结构体中字段的数量
	Field(i int)方法返回字段i的reflect.Value
*/
/*
type order struct {
	ordId      int
	customerId int
}

func createSQL(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField() ; i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createSQL(o)
}
*/

/*
	Int():帮我们去除reflect.Value作为int64
	string():帮我们去除reflect.Value作为string
*/
/*
func main() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}
*/

/*
	升级程序可以变的更加通用，可以适用于任何结构体类型。
	完整程序
*/
type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createSQL(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s%s", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, %s", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createSQL(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createSQL(e)
	i := 90
	createSQL(i)
}
