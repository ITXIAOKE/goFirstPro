package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}

	left, right := 0, len(nums) //左边界
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left > len(nums)-1 || nums[left] > target {
		return res
	}
	res[0] = left

	left, right = 0, len(nums) //右边界
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	res[1] = right - 1
	return res
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 5))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))

}
