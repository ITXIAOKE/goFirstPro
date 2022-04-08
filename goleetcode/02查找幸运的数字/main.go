package main

import "fmt"

//[2,2,3,4] 幸运数字是2
//[2,2,3,3，3] 幸运数字是3，3出现3次，最多
//[2,2,2,3，3] 没有幸运数字，返回-1，因为出现最多的是2,2出现3次，

func findLucky(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	fmt.Println(dic)

	res := -1
	for k, v := range dic  {
		if k == v && v > res {
			res = v
		}
	}
	return res
}

func main() {
	arr := []int{2, 2, 4, 4}
	fmt.Printf("运行数字是：%d", findLucky(arr))
}
