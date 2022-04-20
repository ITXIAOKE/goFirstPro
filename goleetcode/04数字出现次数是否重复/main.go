package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	fmt.Println(dic)

	for _, v := range dic { //这个字典遍历没有顺序123随机显示出来
		//fmt.Println(v)
		if v <= 1 {
			continue
		} else {
			return true
		}

	}

	return false
}

func main() {
	arr := []int{1,3,6,8,9,80}
	fmt.Println(uniqueOccurrences(arr))
}
