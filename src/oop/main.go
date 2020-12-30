package main

import "oop/employee"

/**
go不支持类，而是提供了结构体。结构体中可以添加方法。这样可以将数据和操作数据的方法绑定在一起，实现与类相似的效果。
*/

//func main() {
//	e := employee.Employee {
//		FirstName: "hhy",
//		LastName: "shuai",
//		TotalLeaves: 50,
//		LeavesTaken: 20,
//	}
//	e.LeavesRemaining()
//}

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
