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