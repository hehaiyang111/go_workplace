package main

import "fmt"

/**
   	字符串
 */
func printBytes(s string)  {
	for index,value := range s {
		fmt.Printf("%c starts at byte %d\n", value, index)
	}
}

func printChars(s string) {
	// 用runes防止编码问题
	//runes := []rune(s)
	//for i:= 0; i < len(runes); i++ {
	//	fmt.Printf("%c ",runes[i])
	//}
	for index,value := range s {
		fmt.Printf("%c starts at byte %d\n", value, index)
	}
}

func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
