package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if target >= nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{5, 7, 7, 8, 9, 10}, 8))
	fmt.Println(searchInsert([]int{5, 7, 7, 8, 9, 10}, 6))
	fmt.Println(searchInsert([]int{5, 7, 7, 8, 8, 10}, 0))
	fmt.Println(searchInsert([]int{5, 7, 7, 8, 8, 10}, 11))

}
