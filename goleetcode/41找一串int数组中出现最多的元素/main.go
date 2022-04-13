package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	threshold := len(nums) % 2 //这个只能是取余数，不能是/相除
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val, ok := m[nums[i]] //循环遍历每一遍的元素
		if ok {
			val++ //出现次数+1
			if val >= threshold {
				return nums[i]
			} else {
				m[nums[i]] = val
			}
		} else {
			m[nums[i]] = 1
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{2, 2, 3, 3, 4, 5, 6, 2, 2, 2}))
	fmt.Println(majorityElement([]int{4, 4}))
	fmt.Println(majorityElement([]int{5}))
	fmt.Println(majorityElement([]int{5, 5, 5, 5, 5, 1, 2, 3, 4, 6}))
}
