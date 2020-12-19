package main

import "fmt"

//func main()  {
//	a := [5]int{76, 77, 78, 79, 80}
//	b := a[1:4]
//	fmt.Println(a)
//	fmt.Println(b)
//	//切片的改变  ---->  数组的值也改变
//	for d:= range b {
//		b[d] ++
//	}
//
//	fmt.Print(a)
//}


//func main() {
//	veggies := []string{"potatoes", "tomatoes", "brinjal"}
//	fruits := []string{"oranges", "apples"}
//	// 可变参数一次性传递
//	food := append(veggies, fruits...)
//	fmt.Println("food:",food)
//}


/**
	使用切片的时候，所关联的数组不能删组。对于内存优化是一个问题，解决方法是 copy一下生成一个副本。
 */
func countries() [] string {
	a := []string{"USE","sss","ASADAS","frsa","eqfgfs","ASDsda"}
	b:= a[:len(a) - 2]

	// 复制b切片 创建副本
	c := make([]string , len(b))
	copy(c,b)
	return c
}

func main()  {
	d := countries()
	fmt.Print(d)
}
