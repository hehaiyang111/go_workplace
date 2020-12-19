package main

import "fmt"

/**
	寻找没有重复元素的最长子串
	Example 1:
	Input: "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

	Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

	Input: "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

//func lengthOfLongestSubstring_(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	var freq [256]int
//	result, left, right := 0, 0, 0
//
//	for left < len(s) {
//		// s[right]-'a' 相对位置：对于字符串a是第一个
//		if right < len(s) && freq[s[right]-'a'] == 0 {
//			freq[s[right]-'a']++
//			right++
//		} else {
//			freq[s[left]-'a']--
//			left++
//		}
//		result = max(result, right-left)
//	}
//	return result
//}

func lengthOfLongestSubstring_(a string) int  {
	if 0 == len(a){
		return 0
	}
	var freq [256]int
	result, left , right := 0, 0 ,0
	for left < len(a) {
		if right < len(a) && freq[a[right] - 'a'] == 0{
			freq[a[right] - 'a'] ++
			right ++
		}else {
			freq[a[left] - 'a'] --
			left ++
		}
		result = max(result,right-left)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	a := "bbbbbfdasf"
	fmt.Print(lengthOfLongestSubstring_(a))
}
