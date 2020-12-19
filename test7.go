package main

import "fmt"
/**
	可变参数
	寻找可变参数列表中，有无指定数据。
 */

func find(num int, nums ...int)  {
	fmt.Printf("type of nums is %T\n",nums)
	found := false
	for i,v:=range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main()  {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	//如果直接传入切片
	//nums := []int{89,90,95}
	//find(95,nums...)
}


/**
	append 不会直接改变数组
 */
//func change(s ...string) {
//	s[0] = "Go"
//	// append 会重新创建数组 并复制
//	s = append(s, "playground")
//	fmt.Println(s)
//}
//
//func main() {
//	welcome := []string{"hello", "world"}
//	change(welcome...)
//	fmt.Println(welcome)
//}