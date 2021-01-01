package main

import "fmt"

/**
所有实现了接口的类型，都可以把他的值保存在一个接口类型的变量中。在Go中，我们使用接口的这种特性来实现多态
Example：通过一个程序来理解多态，他会计算一个组织机构的净收益。假设收入来源于两个项目 fixed billing and time and material
*/
// 定义一个接口
type Income interface {
	calculate() int // 计算并返回项目的收入
	source() string // 返回项目的名称
}

type FixedBilling struct {
	projectName  string
	biddedAmount int //表示组织向该项目投资的金额
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

/**
此处是假设，公司出现了新的业务方向可以增加收益，我们添加就非常简单，不需要对calculateNetIncome函数进行任何修改
假设增加了广告收入增益流：
*/
//type Advertisement struct {
//	adName     string
//	CPC        int  //每次点击成本
//	noOfClicks int	//点击次数
//}

// 定义给这些结构体类型定义方法，计算并返回实际收入和项目名称
// FixedBilling和TimeAndMaterial两个结构体都定义了Income的两个方法，因此这两个结构体都实现了Income接口
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

/**
此处是如果增加的广告收益流：
*/
//func (ad Advertisement)source() string {
//	return ad.adName
//}
//
//func (ad Advertisement)calculate() int {
//	return ad.noOfClicks * ad.CPC
//}

/**
目的：接下来声明一个calculateNetIncome函数，用来计算并打印总收入。
ic []Income:接受一个Income接口类型的切片作为参数。该函数会遍历这个切片，并且依个调用calculate()方法，计算出总收入。并且也可以打印收入来源。
多态：根据Income的具体类型我们在calculateNetIncome函数中，程序会调用不同的calculate()and source()方法，如果在该组织以后增加了收入来源,calculateNetIncome则无需修改方法
*/
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	//新的增益流放进来
	//bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	//popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3}
	//新的incomesStreams
	//incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
