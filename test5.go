package main

import (
    "fmt"
)

func main() {
	finger := 8
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: // 默认情况
        fmt.Println("incorrect finger number")
    }
}
/**
	case中可以有多个表达式
**/
func main() {
    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": // 一个选项多个表达式
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}

/**
	在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
**/
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