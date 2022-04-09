package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := []int{-1, -1}

	lookup := make(map[int]int)
	for i, num := range nums {
		temp := target - num           //9-7=2
		if _, ok := lookup[temp]; ok { //查找lookup[2]在map中是否存在，如果存在就把对应的值，赋值给第一个
			res[0] = lookup[temp]
			res[1] = i //把这一轮次的轮次给第二个
			return res
		}
		lookup[num] = i //每一轮都把数字放到这个map里面
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{12, 2, 7, 33, 12}, 9))
}
