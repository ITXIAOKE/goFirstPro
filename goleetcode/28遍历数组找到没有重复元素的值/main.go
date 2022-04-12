package main

import "fmt"

func singleNumber(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	dic := make(map[int]int)
	for _, v := range nums {
		dic[v]++
	}
	num := make([]int, 0)
	for k, v := range dic {
		if v == 1 {
			num = append(num, k)
		}
	}
	return num
}

func main() {
	fmt.Println(singleNumber([]int{1, 12, 12, 3}))
}
