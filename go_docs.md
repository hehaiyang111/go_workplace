# Go

## HelloWorld

```go
package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!")
	fmt.Println(string(1 + '0'))
}
```

## 变量

* 赋值

```go
package main
import "fmt"

func main() {
	 a:=5
	// var a = 4
	fmt.Print(a)
}
```

## 方法

* go的传参是 变量名称 类型名称

* func 方法名称(变量名称1 类型名称1, 变量名称2 类型名称2) 返回值 {}

```go
package main

import (
	"fmt"
)

func calCommodityTotal(price int,num int) int {
	var total = price * num
	return total
}

func main() {
	price,num := 90,100
	totalPrice := calCommodityTotal(price,num)
	fmt.Print("Total price is ", totalPrice)
}
```

## 循环语句

### if

```go
package main

import (  
    "fmt"
)

func main() {  
	num := 99  
    if num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {       // 注意小括号
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }
}
```

* 判断条件没有括号

### for

```go
package main
import (
	"fmt"
)

func main() {
	for i:=1;i<=10;i++ {
		fmt.Printf(" %d",i)
	}
}
```

* for小练习

```go
package main
import(
	"fmt"
)
func main() {
	 //在上面的程序中 no 和 i 被声明然后分别被初始化为 10 和 1 。在每一次循环结束后 no 和 i 都自增 1 。布尔型操作符 && 被用来确保 i 小于等于 10 并且 no 小于等于 19 。
	for a,b := 1,10;a<=10 && b<=19;a,b = a+1,b+1 {
		fmt.Printf("%d * %d = %d\n",b,a,a*b)
	}
}

/*
10 * 1 = 10
11 * 2 = 22
12 * 3 = 36
13 * 4 = 52
14 * 5 = 70
15 * 6 = 90
16 * 7 = 112
17 * 8 = 136
18 * 9 = 162
19 * 10 = 190
 */
```

### switch

```go
func main() {
    finger := 8
    switch finger {
    case 1:
        fmt.Println(Thumb)
    case 2:
        fmt.Println("Index")
    default:
        fmt.Println("incorrect finger number")

    }
    
}
```

* case中可以有多个表达式

```go
func main() {
    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": // 一个选项多个表达式
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}
```

* ​	在 Go 中，**每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行**。使用 **fallthrough **语句可以在**已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中**。

```go
func main() {
	num := 75
    switch num { // num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough    // fallthrough 放在语句后，可以调到下一个 case 强制执行 下一个case
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
```

## 数组

```go
// func number(参数名 参数类型) 返回值的类型 
func number(num [4]int) [4]int {
    num[0] = 778
    return num
}
func main() {
    // [...]可变数组
    num := [...]int{5, 6, 7, 8}
    fmt.Println("传之前的数组：",num)
	fmt.Println(number(num))
	fmt.Println("传之后的数组：",num)
}
```

* 遍历数组

```go
func main() {
    a := [...]float64{67.4, 89.8, 21, 78}
    sum := float64(0)
    for _, v := range a {
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
        fmt.Println("\nsum of all elements of a",sum)

}
```

