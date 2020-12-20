package main

import "fmt"

/**
方法
使用值传递器 在方法内部可见，外部看不到修改
使用指针传递器 外部可以看到
*/

type person struct {
	name string
	age  int
}

type animal struct {
	name string
	age  int
}

func (a person) updateName(newName string) {
	a.name = newName
	fmt.Println("====内部====")
	fmt.Println(a.name)
}

func (b *animal) updateName(newName string) {
	b.name = newName
}

func main() {
	c := person{
		"hhy",
		18,
	}

	d := animal{
		"monkey",
		118,
	}
	fmt.Println("=====人======")
	c.updateName("lxy")
	fmt.Println(c.name)
	fmt.Println("=====动物======")
	d.updateName("pig")
	fmt.Println(d.name)

}
