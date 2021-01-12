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

