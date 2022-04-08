package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{2, 3, 1, 4, 12, 34, 52, 340}))
}

//func QuickSort(ints []int) interface{} {
//
//}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	low := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	high := make([]int, 0, 0)
	flag := nums[0]
	mid = append(mid, flag)

	for i := 1; i < len(nums); i++ { //一定要从第1个开始
		if nums[i] < flag {
			low = append(low, nums[i])
		} else if nums[i] > flag {
			high = append(high, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}
