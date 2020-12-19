package main

import "fmt"

/**
	map 系列
 */

/**
	定义map 并 初始化
 */
//personSalary := make(map[string]int)
//personSalary := map[string]int {
//"steve": 12000,
//"jamie": 15000,
//}

/**
	上面就是获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
*/
//value, ok := map[key]

/**
	判断两个map是否相等
 */
func mapIsEqual(a map[string]int, b map[string]int)  {
	isEqual := false
	if len(a) == len(b) {
		// 遍历其中一个数组
		for key1, value1 := range a {
			// 看另一个map有没有这个key
			value,ok := b[key1]
			// 如果有 判断value, 没有就break掉
			if ok {
				// value相等就ture，继续循环查询
				if value1 == value {
					isEqual = true
				} else {
					//不相等 直接 break掉
					isEqual = false
					break
				}
			}else {
				isEqual = false
				break
			}
		}
	}else {
		isEqual = false
	}

	if isEqual {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}
}

func main()  {
	g := map[string]int {"ss":4,"dd":5,"ww":7}
	b := map[string]int {"ss":4,"dd":5,"ww":7}
	mapIsEqual(g,b)

}
