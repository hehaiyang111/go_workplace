package main
import(
	"fmt"
)
func number(num [4]int) [4]int {
	num[0] = 778
	return num
}
func main() {
	num := [...]int{5,6,7,8}
	fmt.Println("传之前的数组：",num)
	fmt.Println(number(num))
	fmt.Println("传之后的数组：",num)
}

//遍历数组
func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for _, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}