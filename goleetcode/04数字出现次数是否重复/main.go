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

	countDic := make(map[int]int)
	for _, v := range dic {
		_, ok := countDic[v]
		if !ok {
			countDic[v]++
		} else {
			return false
		}
	}
	fmt.Println(countDic)
	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}
