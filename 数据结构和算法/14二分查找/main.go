package main

import "fmt"

/**
二分查找法，在一个【有序的数列】里，把中间元素mid与查找元素target相比较：
		若相等则返回
		若大于target则在mid的右端继续使用二分查找法
		若小于target则在mid的左端继续使用二分查找法
*/
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search(0, len(arr), arr, 5))

	arr2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search2(arr2, 4))
}

//递归方法

func binary_search(left int, right int, arr []int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binary_search(left, right-1, arr, target)
	} else {
		return binary_search(left+1, right, arr, target)
	}
}

func binary_search2(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
