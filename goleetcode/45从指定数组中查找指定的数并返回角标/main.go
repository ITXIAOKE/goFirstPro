package main

import "fmt"

func searchIndex(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target >= nums[0] {
			if nums[mid] < target && nums[0] <= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if target < nums[0] {
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] >= nums[0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(searchIndex([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchIndex([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchIndex([]int{5, 7, 7, 8, 8, 10}, 5))
	fmt.Println(searchIndex([]int{5, 7, 7, 8, 8, 10}, 10))

}
