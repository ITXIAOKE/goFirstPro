package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	if nums == nil || len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				temp := make([]int, 3)
				temp[0] = nums[i]
				temp[1] = nums[left]
				temp[2] = nums[right]
				res = append(res, temp)
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left++
				right -= 1
			} else if sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 2, 2, 1, -1, 0, 0, 3, 6, 4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, 4}))
}
