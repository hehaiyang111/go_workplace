package main

import "fmt"

/**
通过嵌套结构体 可以实现组合
*/
//type author struct {
//	firstName string
//	lastName  string
//	bio       string
//}
//
//func (a author) fullName() string {
//	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
//}
//
//type post struct {
//	title     string
//	content   string
//	author
//}
//
//func (p post) details() {
//	fmt.Println("Title: ", p.title)
//	fmt.Println("Content: ", p.content)
//	fmt.Println("Author: ", p.author.fullName())
//	fmt.Println("Bio: ", p.author.bio)
//}
//
//
//func main() {
//	author1 := author{
//		"Naveen",
//		"Ramanathan",
//		"Golang Enthusiast",
//	}
//	post1 := post{
//		"Inheritance in Go",
//		"Go supports composition instead of inheritance",
//		author1,
//	}
//	post1.details()
//}

/**
使用博客帖子的切片来创建一个网站
*/
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	posts []post //下面需要嵌套结构体，必须起名字，结构体嵌套不能匿名
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	a := author{
		firstName: "hhy",
		lastName:  "zuishuai",
		bio:       "GoLang",
	}
	p1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		a,
	}
	p2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		a,
	}
	p3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		a,
	}
	w := website{
		posts: []post{p1, p2, p3},
	}
	w.contents()
}
