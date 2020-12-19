package main
import(
	"fmt"
)

func printTwoDim(a [3][2]string) {
	// := 赋值并初始化
	for _,v1:=range a {
		for _,v2:= range v1 {
			fmt.Print(v2 + " ")
		}
		fmt.Printf("\n")
	}
}


func main() {
	a:=[3][2]string{
		{"a","b"},
		{"c","d"},
		{"e","f"},    // , 必须有
	}
	printTwoDim(a)
	//fmt.Print(a)
	var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printTwoDim(b)
}

