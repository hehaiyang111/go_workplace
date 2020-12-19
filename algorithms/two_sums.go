package main

import "fmt"

/**
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]
*/
func twoSum(nums []int, target int) []int {
	//利用map实现
	aMap := make(map[int]int)
	//遍历数组
	for i:=0;i<len(nums);i++ {
		another := target - nums[i]
		_,ok:= aMap[another]
		if ok {
			return []int{aMap[another],i}
		}
		aMap[nums[i]] = i
	}
	return nil
}


func main () {
	nums:=[]int{2,7,11,15}
	target:=9
	fmt.Print(twoSum(nums,target))
}
