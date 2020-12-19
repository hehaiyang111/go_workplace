package main
import(
	"fmt"
)
func main() {
	# 在上面的程序中 no 和 i 被声明然后分别被初始化为 10 和 1 。在每一次循环结束后 no 和 i 都自增 1 。布尔型操作符 && 被用来确保 i 小于等于 10 并且 no 小于等于 19 。
	for a,b := 1,10;a<=10 && b<=19;a,b = a+1,b+1 {
		fmt.Printf("%d * %d = %d\n",b,a,a*b)
	}
}