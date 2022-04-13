package main

import (
	"fmt"
)

//删除指定的元素后，求剩下的元素

func removeElementAfter(nums []int, val int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	r := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			r = append(r, nums[i])
		}
	}
	return r
}

//删除指定的元素后，求剩下元素的个数

func removeElementAfterN(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	r := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			r += 1
		}
	}
	return r
}

func main() {
	fmt.Println(removeElementAfter([]int{-1, 2, 2, 1, -1, 0, 0, 3, 6, 4}, 2))
	fmt.Println(removeElementAfterN([]int{-1, 2, 2, 1, -1, 0, 0, 3, 6, 4}, 2))

	fmt.Println(removeElementAfter([]int{3, 2, 2, 3}, 2))
	fmt.Println(removeElementAfterN([]int{3, 2, 2, 3}, 2))
}
